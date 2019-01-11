package acct

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqNuLoginPersona struct {
	// TXN=NuLoginPersona
	Txn  string `fesl:"TXN"`
	Name string `fesl:"name"` // Value specified in +soldierName
}

type ansNuLoginPersona struct {
	Txn       string `fesl:"TXN"`
	ProfileID int    `fesl:"profileId"`
	UserID    int    `fesl:"userId"`
	LobbyKey  string `fesl:"lkey"`
}

// NuLoginPersona handles acct.NuLoginPersona command
func (acct *Account) NuLoginPersona(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuLoginPersona(event)
	default:
		acct.clientNuLoginPersona(event)
	}
}

// clientNuLoginPersona used when user selects hero from game client
func (acct *Account) clientNuLoginPersona(event network.EventClientCommand) {

	h := 1

	event.Client.PlayerData.HeroID = h
	event.Client.PlayerData.PlayerID = h
	event.Client.PlayerData.HeroName = "1234"

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginPersona{
			Txn:       acctNuLoginPersona,
			ProfileID: event.Client.PlayerData.PlayerID,
			UserID:    event.Client.PlayerData.HeroID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

// NuLoginPersonaServer - soldier login command
func (acct *Account) serverNuLoginPersona(event network.EventClientCommand) {
	a := 1

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginPersona{
			Txn:       acctNuLoginPersona,
			ProfileID: a,
			UserID:    a,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}
