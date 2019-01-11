package network

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

func (client *Client) readFESL(data []byte) {
	cmds, err := codec.ParseCommands(data)
	if err != nil {
		logrus.
			WithError(err).
			WithField("packet", string(data)).
			Error("Cannot parse commands")
		return
	}
	for _, cmd := range cmds {
		client.receiver <- ClientEvent{
			Name: "command." + cmd.Query,
			Data: cmd,
		}
	}
}

func (client *Client) readTLSPacket(data []byte) {
	cmds, err := codec.ParseCommands(data)
	if err != nil {
		var extract string
		if len(data) > 128 {
			extract = string(data[:128])
		} else {
			extract = string(data)
		}

		logrus.
			WithError(err).
			WithField("extract", extract).
			Error("Cannot parse commands (TLS)")
		return
	}
	for _, cmd := range cmds {
		client.receiver <- ClientEvent{
			Name: "command." + cmd.Message["TXN"],
			Data: cmd,
		}
	}
}
