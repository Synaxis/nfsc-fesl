package theater

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
	"github.com/Synaxis/nfsc-fesl/backend/ranking"
)

// EGAM is sent to Game-Client
type reqEGAM struct {
	GameID       int    `fesl:"GID"`
	LobbyID      int    `fesl:"LID"`
	Port         int    `fesl:"PORT"`
	PlatformType int    `fesl:"PTYPE"`
	RemoteIP     string `fesl:"R-INT-IP"`
	RemotePort   int    `fesl:"R-INT-PORT"`
	AccountID    int    `fesl:"R-U-accid"`    //RUAccid:  stats["userID"]
	Category     int    `fesl:"R-U-category"` // TODO: What exactly it is?
	Region       string `fesl:"R-U-dataCenter"`
	StatsElo     int    `fesl:"R-U-elo"`
	ExternalIP   string `fesl:"R-U-externalIp"`
	StatsKit     int    `fesl:"R-U-kit"`
	StatsLevel   int    `fesl:"R-U-lvl"`
	StatsTeam    int    `fesl:"R-U-team"`
	TID          int    `fesl:"TID"`
}

type ansEGAM struct {
	TID     string `fesl:"TID"`
	LobbyID string `fesl:"LID"`
	GameID  int    `fesl:"GID"`
}

// a EGAM - CLIENT called when a client wants to join a gameServer
func (tm *Theater) EnterGame(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse value of GID in theater.EGAM")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if gameID < 1 {
		logrus.Println("cannot find gameID")
	}
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Not found any server when joining game")
		return
	}

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGame,
		Payload: ansEGAM{
			event.Command.Message["TID"],
			event.Command.Message["LID"],
			gameID,
		},
	})

	gr := GameRequest{
		PlayerID: 1,
		HeroID:   1,
		HeroName: "1234",
		GameID:   1,
		LobbyID:  event.Command.Message["LID"],
		Ticket:   "2018751182",
	}

	tm.EnterGameRequest(&event, game.GameServer, gr)
	tm.EGEG(&event, game.GameServer, gr)
}

type GameRequest struct {
	PlayerID int
	HeroID   int
	HeroName string
	GameID   int
	LobbyID  string
	Stats    ranking.Stats
	Ticket   string
}
