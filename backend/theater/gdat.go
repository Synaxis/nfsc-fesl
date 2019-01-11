package theater

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqGDAT struct {
	// TID=3
	TID int `fesl:"TID"`

	// LID=0
	LobbyID int `fesl:"LID"`
	// GID=1
	GameID int `fesl:"GID"`
}

type ansGDAT struct {
	TID string `fesl:"TID"`

	EloRank             string `fesl:"B-U-elo_rank"`
	AvgAllyRank         string `fesl:"B-U-avg_ally_rank"`
	AvgAxisRank         string `fesl:"B-U-avg_axis_rank"`
	ArmyDistribution    string `fesl:"B-U-army_distribution"`
	ArmyBalance         string `fesl:"B-U-army_balance"`
	PercentFull         string `fesl:"B-U-percent_full"`
	AvailSlotsNational  string `fesl:"B-U-avail_slots_national"`
	AvailSlotsRoyal     string `fesl:"B-U-avail_slots_royal"`
	AvailableVipsNation string `fesl:"B-U-avail_vips_national"`
	AvailableVipsRoyal  string `fesl:"B-U-avail_vips_royal"`
	IsRanked            string `fesl:"B-U-ranked"`
	Easyzone            string `fesl:"B-U-easyzone"`
	ServerType          string `fesl:"B-U-servertype"`
	ServerState         string `fesl:"B-U-server_state"`
	MapName             string `fesl:"B-U-map_name"`
	PunkBusterEnabled   string `fesl:"B-U-punkb"`
	StdDevLevel         string `fesl:"B-U-lvl_sdv"`
	AvgLevel            string `fesl:"B-U-lvl_avg"`

	GameID     int    `fesl:"GID"`
	Join       string `fesl:"JOIN"`
	ServerName string `fesl:"NAME"`

	AP      string `fesl:"AP"` // PlayerTypeCount, int
	LobbyID int    `fesl:"LID"`

	// O / W / C (default=O, but if there will some random string not eaual to W and C it will also work)
}

// GDAT - CLIENT called to get data about the server
func (tm *Theater) GameData(event network.EventClientCommand) {
	gameID, err := event.Command.Message.IntVal("GID")
	if err != nil {
		logrus.WithError(err).Warn("Cannot parse GID in theater.GDAT")
		return
	}

	game, err := tm.mm.GetGame(gameID)
	if err != nil {
		logrus.
			WithError(err).
			WithField("gameID", gameID).
			Warn("Camtt find Game in matchmaking")
		return
	}
	msg := game.GameServer.ServerData

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrGamesData,
		Payload: ansGDAT{
			LobbyID:             game.LobbyID,
			AP:                  game.GameServer.ServerData.Get("AP"),
			TID:                 event.Command.Message["TID"],
			GameID:              game.ID,
			Join:                msg.Get("JOIN"),
			ServerName:          msg.Get("NAME"),
			EloRank:             msg.Get("B-U-elo_rank"),
			AvgAllyRank:         msg.Get("B-U-avg_ally_rank"),
			AvgAxisRank:         msg.Get("B-U-avg_axis_rank"),
			ArmyDistribution:    msg.Get("B-U-army_distribution"),
			ArmyBalance:         msg.Get("B-U-army_balance"),
			PercentFull:         msg.Get("B-U-percent_full"),
			AvailSlotsNational:  msg.Get("B-U-avail_slots_national"),
			AvailSlotsRoyal:     msg.Get("B-U-avail_slots_royal"),
			AvailableVipsNation: msg.Get("B-U-avail_vips_national"),
			AvailableVipsRoyal:  msg.Get("B-U-avail_vips_royal"),
			IsRanked:            msg.Get("B-U-ranked"),
			Easyzone:            msg.Get("B-U-easyzone"),
			ServerType:          msg.Get("B-U-servertype"),
			ServerState:         msg.Get("B-U-server_state"),
			PunkBusterEnabled:   msg.Get("B-U-punkb"),
			MapName:             msg.Get("B-U-map_name"),
			AvgLevel:            msg.Get("B-U-lvl_avg"),
			StdDevLevel:         msg.Get("B-U-lvl_sdv"),
		},
	})
}
