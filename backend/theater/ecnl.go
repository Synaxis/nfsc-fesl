package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqECNL struct {
	// TID=8
	TID int `fesl:"TID"`
	// GID=3
	GameID int `fesl:"GID"`
	// LID=1
	LobbyID int `fesl:"LID"`
}

type ansECNL struct {
	TID     string `fesl:"TID"`
	GameID  string `fesl:"GID"`
	LobbyID string `fesl:"LID"`
}

func (tm *Theater) EnterConnectionLAN(event network.EventClientCommand) {
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterConnectionLost,
		Payload: ansECNL{
			event.Command.Message["TID"],
			event.Command.Message["GID"],
			event.Command.Message["LID"],
		},
	})
}
