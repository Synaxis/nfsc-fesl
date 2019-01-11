package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqPLVT struct {
	// TID=16
	TID int `fesl:"TID"`

	// LID=1
	LobbyID int `fesl:"LID"`
	// GID=12
	GameID int `fesl:"GID"`
	// PID=3
	PlayerID int `fesl:"PID"`
}

type ansKICK struct {
	PlayerID string `fesl:"PID"`
	LobbyID  string `fesl:"LID"`
	GameID   string `fesl:"GID"`
}

type ansPLVT struct {
	TID string `fesl:"TID"`
}

func (tm *Theater) PlayerExited(event network.EventClientCommand) {
	// pid := event.Command.Message["PID"]
	// if j, err := json.Marshal(msgTeamGid{event.Command.Message["GID"], stats["c_team"]}); err == nil {
	// 	tm.pub.Publish(queue.TopicTeamLeft, j)
	// }

	event.Client.WriteEncode(&codec.Answer{
		Type: thtrKICK,
		Payload: ansKICK{
			event.Command.Message["PID"],
			event.Command.Message["LID"],
			event.Command.Message["GID"],
		},
	})

	event.Client.WriteEncode(&codec.Answer{
		Type:    thtrPLVT,
		Payload: ansPLVT{event.Command.Message["TID"]},
	})
}
