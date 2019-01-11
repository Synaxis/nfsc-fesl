package pnow

import (
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

const (
	sessionStateNEW       = "NEW"
	sessionStateDeleted   = "DELETED"
	sessionStateCancelled = "CANCELLED"
	sessionStateError     = "ERROR"
	sessionStateComplete  = "COMPLETE"
	sessionStateActive    = "ACTIVE"
)

const (
	resultTypeJoin = "JOIN"
)

type ansStatus struct {
	Txn          string                 `fesl:"TXN"`
	ID           statusPartition        `fesl:"id"`
	SessionState string                 `fesl:"sessionState"`
	Properties   map[string]interface{} `fesl:"props"`
}

type statusPartition struct {
	ID        int    `fesl:"id"`
	Partition string `fesl:"partition"`
}

type statusGame struct {
	LobbyID   int `fesl:"lid"`
	FitFactor int `fesl:"fit"`
	GameID    int `fesl:"gid"`
}

// Status handles pnow.Status command
func (pnow *PlayNow) Status(event network.EventClientCommand) {
	// TODO: Use matchmaking.Pool to prepare list of available games
	game, err := pnow.MM.FindAvailableGame()
	if err != nil {
		logrus.WithError(err).Warn("Cannot list available games in pnow.Status")
		// Return an error to client about failed matchmaking
		pnow.answer(event.Client, 0x80000000, ansStatus{
			Txn:          pnowStatus,
			ID:           statusPartition{1, event.Command.Message["partition.partition"]},
			SessionState: sessionStateError,
		})
		return
	}

	games := []statusGame{
		{
			LobbyID:   game.LobbyID,
			GameID:    game.ID,
			FitFactor: 1001,
		},
	}

	pnow.answer(
		event.Client,
		0x80000000, // event.Command.PayloadID,
		ansStatus{
			Txn:          pnowStatus,
			ID:           statusPartition{1, event.Command.Message["partition.partition"]},
			SessionState: sessionStateComplete,
			Properties: map[string]interface{}{
				"resultType": resultTypeJoin,
				"games":      games,
			},
		},
	)
}
