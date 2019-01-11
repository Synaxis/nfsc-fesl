package network

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/config"
)

// Socket is a basic event-based TCP-Server
// TODO: Rename it to broker
type Socket struct {
	bind      string
	listen    net.Listener
	EventChan chan SocketEvent
}

func newSocket(bind string) *Socket {
	return &Socket{
		bind:      bind,
		EventChan: make(chan SocketEvent),
	}
}

func NewSocketTCP(bind string) (*Socket, error) {
	socket := newSocket(bind)
	listener, err := socket.listenTCP()
	if err != nil {
		return nil, err
	}
	socket.listen = listener
	go socket.run(socket.createClientTCP)

	return socket, nil
}

func NewSocketTLS(bind string) (*Socket, error) {
	socket := newSocket(bind)
	listener, err := socket.listenTLS()
	if err != nil {
		return nil, err
	}
	socket.listen = listener
	go socket.run(socket.createClientTLS)

	return socket, nil
}

func (socket *Socket) listenTCP() (net.Listener, error) {
	listener, err := net.Listen("tcp", socket.bind)
	if err != nil {
		logrus.WithError(err).Errorf("Listening on %s threw an error", socket.bind)
		return nil, err
	}
	return listener, nil
}

func (socket *Socket) listenTLS() (net.Listener, error) {
	cert, err := config.ParseCertificate()
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ClientAuth:         tls.NoClientCert,
		MinVersion:         tls.VersionSSL30,
		InsecureSkipVerify: true,
		//MaxVersion:   tls.VersionSSL30,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_RC4_128_SHA,
		},
	}

	listener, err := tls.Listen("tcp", socket.bind, config)
	if err != nil {
		logrus.WithError(err).Errorf("Listening on %s threw an error", socket.bind)
		return nil, err
	}

	return listener, nil
}

type connAcceptFunc func(conn net.Conn)

func (socket *Socket) run(connect connAcceptFunc) {
	for {
		// Wait and listen for incomming connection
		conn, err := socket.listen.Accept()
		if err != nil {
			logrus.WithError(err).Errorf("A new client connecting to %s threw an error", socket.bind)
			continue
		}

		// Establish connection
		connect(conn)
	}
}

func (socket *Socket) createClientTCP(conn net.Conn) {
	tcpClient := NewClientTCP(conn)
	go tcpClient.handleRequestTCP()
	go tcpClient.handleClientEvents(socket)
	socket.EventChan <- socket.FireNewClient(tcpClient)
}

func (socket *Socket) createClientTLS(conn net.Conn) {
	tlscon, ok := conn.(*tls.Conn)
	if !ok {
		conn.Close()
		return
	}

	tlscon.SetDeadline(time.Now().Add(time.Second * 10))
	err := tlscon.Handshake()
	if err != nil {
		logrus.WithError(err).Errorf("A new client from %s connecting to %s threw an error", tlscon.RemoteAddr(), socket.bind)
		tlscon.Close()
	}
	tlscon.SetDeadline(time.Time{})

	state := tlscon.ConnectionState()
	logrus.Debugf("Connection handshake complete %t, %v", state.HandshakeComplete, state)

	tlsClient := NewClientTLS(tlscon)
	go tlsClient.handleRequestTLS()
	go tlsClient.handleClientEvents(socket)

	logrus.
		WithField("bind", socket.bind).
		WithField("protocol", "tcp").
		Print("A new client connected")

	socket.EventChan <- socket.FireNewClient(tlsClient)
}

// SocketEvent is the generic struct for events
// by this socket
//
// Current events:
//		Name				-> Data-Type
// 		close 				-> nil
//		error				-> error
//		newClient			-> *Client
//		client.close		-> [0: *client, 1:nil]
//		client.error		-> {*client, error}
//		client.command.*	-> {*client, *Command}
//		client.data			-> {*client, string}
type SocketEvent struct {
	Name string
	Data interface{}
}

func (s *Socket) FireNewClient(c *Client) SocketEvent {
	return SocketEvent{
		Name: "newClient",
		Data: EventClientCommand{Client: c},
	}
}

// Close fires a close-event and closes the socket
func (socket *Socket) Close() {
	// Fire closing event
	close(socket.EventChan)
	

	// Close socket
	socket.listen.Close()
}
