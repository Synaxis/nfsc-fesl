package codec

import (
	"bytes"
	"encoding/binary"

	"github.com/sirupsen/logrus"
)

type Command struct {
	Query     string
	Message   Fields
	PayloadID uint32
}

func NewCommand(pkt *RawPacket) (*Command, error) {
	l := len(pkt.Payload)
	if l > 8096 {
		l = 8096
	}

	logrus.
		WithField("query", string(pkt.Query)).
		WithField("payload", string(pkt.Payload[:l])).
		WithField("broadcast", pkt.Broadcast[0]).
		WithField("id", pkt.Broadcast[1:]).
		Debug("codec.NewCommand")

	out := &Command{
		Query:     string(pkt.Query),
		PayloadID: binary.BigEndian.Uint32(pkt.Broadcast),
		Message:   DecodeFESL(pkt.Payload),
	}

	return out, nil
}

func ParseCommands(data []byte) ([]*Command, error) {
	buf := bytes.NewBuffer(data)
	cmds := []*Command{}

	for buf.Len() > 0 {
		pkt, err := ExtractPacket(buf)
		if err != nil {
			return nil, err
		}

		cmd, err := NewCommand(pkt)
		if err != nil {
			return nil, err
		}

		cmds = append(cmds, cmd)
	}

	return cmds, nil
}
