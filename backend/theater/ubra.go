package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqUBRA struct {
	// TID=5
	TID int `fesl:"TID"`

	// LID=1
	LobbyID int `fesl:"LID"`
	// GID=3
	GameID int `fesl:"GID"`
	// START=0
	// START=1
	START int `fesl:"START"`
}

type ansUBRA struct {
	TID string `fesl:"TID"`
}

func (tm *Theater) UpdateBracket(event network.EventClientCommand) {
	// Reset AP counter
	if event.Command.Message["START"] == "1" {
		// hash := event.Client.GetServerData(event.Command.Message["GID"])
		event.Client.ServerData.Set("AP", "0")
	}

	event.Client.WriteEncode(&codec.Answer{
		Type:    thtrUBRA,
		Payload: ansUBRA{event.Command.Message["TID"]},
	})
}
