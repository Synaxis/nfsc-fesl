package acct

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqServerNuGetPersonas struct {
	// TXN=NuGetPersonas
	TXN string `fesl:"TXN"`
	// namespace=
	Namespace string `fesl:"namespace"`
}

type ansNuGetPersonas struct {
	Txn      string   `fesl:"TXN"`
	Personas []string `fesl:"personas"`
}

// NuGetPersonas handles acct.NuGetPersonas command
// NuGetPersonas - Soldier data lookup call
func (acct *Account) NuGetPersonas(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuGetPersonas(event)
	default:
		acct.clientNuGetPersonas(event)
	}
}

func (acct *Account) clientNuGetPersonas(event network.EventClientCommand) {

	ans := ansNuGetPersonas{Txn: acctNuGetPersonas, Personas: []string{}}

	acct.answer(event.Client, event.Command.PayloadID, ans)
}

func (acct *Account) serverNuGetPersonas(event network.EventClientCommand) {
	srv := 1

	event.Client.PlayerData.ServerID = srv
	event.Client.PlayerData.ServerUserName = "MargeSimpson"

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuGetPersonas{
			Txn:      acctNuGetPersonas,
			Personas: []string{"MargeSimpson"},
		},
	)
}
