package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqPENT struct {
	// TID=7
	TID int `fesl:"TID"`
	// LID=1
	LobbyID int `fesl:"LID"`
	// GID=72
	GameID int `fesl:"GID"`
	// PID=733
	PlayerID int `fesl:"PID"`
}

type ansPENT struct {
	TID      string `fesl:"TID"`
	PlayerID string `fesl:"PID"`
}

// PlayerEntered player joined game
func (tm *Theater) PlayerEntered(event network.EventClientCommand) {
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrPlayerEnter,
		Payload: ansPENT{
			event.Command.Message["TID"],
			event.Command.Message["PID"],
		},
	})
}
