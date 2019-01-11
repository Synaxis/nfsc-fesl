package network

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type SocketUDPEvent struct {
	Name string
	Addr *net.UDPAddr
	Data interface{}
}

type SocketUDP struct {
	bind      string
	listen    *net.UDPConn
	EventChan chan SocketUDPEvent
}

func NewSocketUDP(bind string) (*SocketUDP, error) {
	socket := &SocketUDP{
		bind:      bind,
		EventChan: make(chan SocketUDPEvent, 1000),
	}

	var err error
	serverAddr, err := net.ResolveUDPAddr("udp", socket.bind)
	if err != nil {
		logrus.WithError(err).Errorf("Listening on %s threw an error", socket.bind)
		return nil, err
	}

	socket.listen, err = net.ListenUDP("udp", serverAddr)
	if err != nil {
		logrus.WithError(err).Errorf("Listening on %s threw an error", socket.bind)
		return nil, err
	}

	go socket.run()

	return socket, nil
}

func (socket *SocketUDP) run() {
	buf := make([]byte, 2048) //test
	//buf := make([]byte, 8096)
	//the write may have to wait until the read is finished. For a one-byte file, this is the best you can do, but for a 1MB file, this will be extremely slow.


	for socket.EventChan != nil {
		n, addr, err := socket.listen.ReadFromUDP(buf)
		if err != nil {
			logrus.WithError(err).Error("Error reading from UDP", err)
			socket.EventChan <- SocketUDPEvent{Name: "error", Addr: addr, Data: err}
			continue
		}

		socket.readFESL(buf[:n], addr)
	}
}

func (socket *SocketUDP) readFESL(data []byte, addr *net.UDPAddr) {
	p := bytes.NewBuffer(data)
	var payloadID uint32
	var payloadLen uint32

	payloadType := string(data[:4])
	p.Next(4)

	binary.Read(p, binary.BigEndian, &payloadID)
	binary.Read(p, binary.BigEndian, &payloadLen)

	payload := codec.DecodeFESL(data[12:])

	socket.EventChan <- SocketUDPEvent{
		Name: payloadType,
		Addr: addr,
		Data: &codec.Command{
			Query:     payloadType,
			PayloadID: payloadID,
			Message:   payload,
		},
	}
}

func (socket *SocketUDP) WriteEncode(packet *codec.Answer, addr *net.UDPAddr) error {
	// Encode packet
	buf, err := codec.
		NewEncoder().
		EncodePacket(packet)
	if err != nil {
		logrus.
			WithError(err).
			WithField("type", packet.Type).
			Error("Cannot encode packet")
		return err
	}

	// Send packet
	_, err = socket.listen.WriteTo(buf.Bytes(), addr)
	if err != nil {
		logrus.
			WithError(err).
			WithField("type", packet.Type).
			Warn("Cannot send encoded packet")
		return err
	}

	return nil
}

// Close fires a close-event and closes the socket
func (socket *SocketUDP) Close() {
	// Fire closing event
	socket.EventChan <- SocketUDPEvent{Name: "close", Addr: nil, Data: nil}

	// Close socket
	socket.listen.Close()
}
