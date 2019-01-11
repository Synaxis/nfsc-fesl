package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type ansECHO struct {
	TID       string `fesl:"TID"`
	Txn       string `fesl:"TXN"`
	IP        string `fesl:"IP"`
	Port      int    `fesl:"PORT"`
	ErrStatus int    `fesl:"ERR"`
	Type      int    `fesl:"TYPE"`
}

func (tm *Theater) ECHO(event network.SocketUDPEvent) {
	command := event.Data.(*codec.Command)

	tm.socketUDP.WriteEncode(&codec.Answer{
		Type: codec.ThtrEcho,
		Payload: ansECHO{
			TID:       command.Message["TID"],
			IP:        event.Addr.IP.String(),
			Port:      event.Addr.Port,
			ErrStatus: 0,
			Type:      1,
		},
	}, event.Addr)
}
