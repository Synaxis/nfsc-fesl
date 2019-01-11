package fsys

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

const (
	fsysGetPingSites = "GetPingSites"
	fsysHello        = "Hello"
	fsysMemCheck     = "MemCheck"
	fsysPing         = "Ping"
)

type ConnectSystem struct {
	ServerMode bool
}

func (fsys *ConnectSystem) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslSystem,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
