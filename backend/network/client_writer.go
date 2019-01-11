package network

import (
	//"errors"

	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

func (c *Client) SendPacket(pkt []byte) error {
	_, err := c.Conn.Write(pkt)
	if err != nil {
		logrus.
			WithError(err).
			Warn("Cannot send encoded packet")
		return err
	}

	logrus.
		WithField("packet", string(pkt)).
		Print("client.SendPacket")

	return nil
}

func (c *Client) WriteEncode(packet *codec.Answer) error {
	// if !c.IsActive {                                       -- > theater is UDP 
	// 	logrus.Warnf("Trying to write to inactive Client.\n%v", packet.Type)     -- > theater is UDP 
	// 	return errors.New("Client is not active. Can't send message")  -- > theater is UDP 
	// }

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

	return c.SendPacket(buf.Bytes())
}
