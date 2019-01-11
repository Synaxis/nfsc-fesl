package acct

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqNuGetAccount struct {
	// TXN=NuGetAccount
	TXN string `fesl:"TXN"`
}

type ansNuGetAccount struct {
	TXN string `fesl:"TXN"`

	DobDay         int    `fesl:"DOBDay"`
	DobMonth       int    `fesl:"DOBMonth"`
	DobYear        int    `fesl:"DOBYear"`
	Country        string `fesl:"country"`
	Language       string `fesl:"language"`
	GlobalOptIn    bool   `fesl:"globalOptin"`
	ThidPartyOptIn bool   `fesl:"thirdPartyOptin"`
}

type ansClientNuGetAccount struct {
	ansNuGetAccount

	NucleusID int    `fesl:"nuid"`
	UserID    int    `fesl:"userId"`
	HeroName  string `fesl:"heroName"`
}

// NuGetAccount handles acct.NuGetAccount command
func (acct *Account) NuGetAccount(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuGetAccount(event)
	default:
		acct.clientNuGetAccount(event)
	}
}

func (acct *Account) clientNuGetAccount(event network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansClientNuGetAccount{
			ansNuGetAccount: ansNuGetAccount{
				TXN:            acctNuGetAccount,
				Country:        "US",
				Language:       "en_US",
				DobDay:         1,
				DobMonth:       1,
				DobYear:        1990,
				GlobalOptIn:    false,
				ThidPartyOptIn: false,
			},

			NucleusID: event.Client.PlayerData.PlayerID,
			UserID:    event.Client.PlayerData.PlayerID,
			HeroName:  event.Client.PlayerData.HeroName, // ?
		},
	)
}

func (acct *Account) serverNuGetAccount(event network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuGetAccount{
			TXN:            acctNuGetAccount,
			Country:        "US",
			Language:       "en_US",
			DobDay:         1,
			DobMonth:       1,
			DobYear:        1990,
			GlobalOptIn:    false,
			ThidPartyOptIn: false,
		},
	)
}
