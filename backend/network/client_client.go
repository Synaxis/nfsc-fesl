package network

import (
	"crypto/tls"
	"io"
	"net"
	"strings"
	"time"
	"sync"
	"fmt"


	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

const (
	// FragmentSize is defined in fsys.Hello which specifies
	// how many bytes can be in one packet sent over the backend.
	//
	// TODO: Use value specified in fsys.Hello instead of this hardcoded const
	FragmentSize = 8096
)

type cliKey struct {
	name, addr string
}
func (ck *cliKey) String() string {
	return fmt.Sprintf("%s_%s", ck.name, ck.addr)
}

type clientsArray struct { //What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
	mu        *sync.Mutex
	connected map[cliKey]*Client
}

func newclientsArray() *clientsArray {
	return &clientsArray{
		connected: make(map[cliKey]*Client, 500),
		mu:        new(sync.Mutex),
	}
}

func (cls *clientsArray) Add(cl *Client) {
	cls.mu.Lock()
	cls.connected[cl.Key()] = cl
	cls.mu.Unlock()
}

func (cls *clientsArray) Remove(cl *Client) {
	cls.mu.Lock()
	delete(cls.connected, cl.Key())
	cls.mu.Unlock()
}

//key
func (c *Client) Key() cliKey {
	return cliKey{c.name, c.IpAddr.String()}
} 



type Client struct {
	name		string
	Conn        net.Conn
	receiver    chan ClientEvent
	sender      chan codec.Answer
	IsActive    bool
	IpAddr      net.Addr
	HeartTicker *time.Ticker

	// Type defines a what type of executable client defines this connection
	// i.e. "server", "client-nonreg"
	Type string

	PlayerData *PlayerData
	ServerData ServerData
}

func newClient(conn net.Conn) *Client {
	return &Client{
		Conn:       conn,
		IpAddr:     conn.RemoteAddr(),
		receiver:   make(chan ClientEvent, 5),
		sender:     make(chan codec.Answer, 5),
		IsActive:   true,
		PlayerData: &PlayerData{},
		ServerData: ServerData{},
	}
}

func NewClientTCP(conn net.Conn) *Client {
	return newClient(conn)
}

func NewClientTLS(conn *tls.Conn) *Client {
	return newClient(conn)
}

func (c *Client) GetClientType() string {
	return c.Type
}

func (c *Client) handleRequestTLS() {
	c.IsActive = true
	buf := make([]byte, FragmentSize)

	for c.IsActive {
		n, err := c.readBuf(buf)
		if err != nil {
			return
		}

		c.readTLSPacket(buf[:n])
		buf = make([]byte, FragmentSize)
	}
}

func (c *Client) handleRequestTCP() {
	c.IsActive = true
	buf := make([]byte, FragmentSize)

	for c.IsActive {
		n, err := c.readBuf(buf)
		if err != nil {
			return
		}

		c.readFESL(buf[:n])
		buf = make([]byte, FragmentSize)
	}
}

func (c *Client) readBuf(buf []byte) (int, error) {
	n, err := c.Conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			logrus.Errorf("Error: %v on client %s", err, c.IpAddr)
			c.receiver <- c.FireClose()
			return 0, err
		}
		c.receiver <- c.FireClose()
		return 0, err
	}
	return n, nil
}

func (c *Client) handleClientEvents(socket *Socket) {
	defer c.Close()

	for c.IsActive {
		select {
		case event := <-c.receiver:
			switch {
			case event.Name == "close":
				return
			case strings.Index(event.Name, "command") != -1:
				socket.EventChan <- c.FireClientCommand(event)
			case event.Name == "data":
				logrus.Warnf("Not implemented: Client send client.data: %s", event.Data)
			default:
				logrus.Warn("Not implemented client.%s for %s", event.Name, event.Data)
			}
		}
	}
}

func (c *Client) Close() {
	c.IsActive = false
	logrus.WithField("ip", c.IpAddr).Print("Client closing connection")
	c.Conn.Close()
	if lkey := c.PlayerData.LobbyKey; lkey != "" {
		Lobby.Delete(lkey)
	}
	close(c.receiver)
	close(c.sender)
}
