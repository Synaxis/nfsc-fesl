package rank

import (
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqGetStats struct {
	Owner      string   `fesl:"owner"`      // owner=5
	OwnerType  string   `fesl:"ownerType"`  // ownerType=1
	PeriodID   string   `fesl:"periodId"`   // periodId=0
	PeriodPast string   `fesl:"periodPast"` // periodPast=0
	Keys       []string `fesl:"keys"`       // keys.0=c_apr, keys.1=level ...
}

type ansGetStats struct {
	Txn       string `fesl:"TXN"`
	OwnerID   int    `fesl:"ownerId"`
	OwnerType int    `fesl:"ownerType"`
	//Stats     []statsPair `fesl:"stats"`
}

type statsPair struct {
	Key   string `fesl:"key"`
	Text  string `fesl:"text,omitempty"`
	Value string `fesl:"value"`
}

// GetStats - Get basic stats about a soldier/owner (account holder)
func (r *Ranking) GetStats(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverGetStats(&event)
	default:
		r.clientGetStats(&event)
	}
}

func (r *Ranking) serverGetStats(event *network.EventClientCommand) {
	r.getStats(event)
}

func (r *Ranking) clientGetStats(event *network.EventClientCommand) {
	r.getStats(event)
}

func (r *Ranking) getStats(event *network.EventClientCommand) {
	if event.Command.Message["owner"] == "Current" {
		// In the tutorial "Current" is reserved name for the hero.
		return
	}

	ownerID, err := event.Command.Message.IntVal("owner")
	if err != nil {
		logrus.
			WithField("owner", event.Command.Message["owner"]).
			WithField("cmd", "rank.GetStats").
			Warnf("Cannot parse ownerID")
		return
	}

	heroID := event.Client.PlayerData.HeroID
	if heroID == 0 {
		if event.Client.Type == "server" {
			heroID = ownerID
			// Server uses only heroID to identify owners
			logrus.Warnf("GetStats (server), replacing heroID with ownerID")
		} else {
			// Whether client is not yet logged in it requires data about the
			// tutorial completion, sadly we need to look for a player's master
			// hero

			heroID = 1
		}
	}

	r.answer(
		event.Client,
		event.Command.PayloadID,
		ansGetStats{
			Txn:       rankGetStats,
			OwnerID:   heroID,
			OwnerType: 1,
			//Stats:     [],
		},
	)
}
