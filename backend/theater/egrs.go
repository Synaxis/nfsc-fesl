package theater

import (
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqEGRS struct {
	// TID=6
	TID int `fesl:"TID"`

	// LID=1
	LobbyID int `fesl:"LID"`
	// GID=12
	GameID int `fesl:"GID"`
	// ALLOWED=1
	// ALLOWED=0
	IsAllowed int `fesl:"ALLOWED"`
	// PID=3
	PlayerID int `fesl:"PID"`

	// Reason is only sent when ALLOWED=0 and there is some kind of error
	// REASON=-602
	Reason string `fesl:"REASON,omitempty"`
}

type ansEGRS struct {
	TID string `fesl:"TID"`
}

// EGRS - SERVER sent up, tell us if client is 'allowed' to join
func (tm *Theater) EnterGameHostResponse(event network.EventClientCommand) {
	if event.Command.Message["ALLOWED"] == "0" {
		logrus.
			WithFields(logrus.Fields{
				"gameID": event.Command.Message["GID"],
				"heroID": event.Command.Message["PID"],
				"reason": event.Command.Message["REASON"],
			}).
			Warn("EGRS: Player cannot join server, look for REASON code")
	}

	event.Client.WriteEncode(&codec.Answer{
		Type:    codec.ThtrEnterGameResponse,
		Payload: ansEGRS{event.Command.Message["TID"]},
	})
}
