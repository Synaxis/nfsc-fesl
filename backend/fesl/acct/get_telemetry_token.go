package acct

import (
	"github.com/Synaxis/nfsc-fesl/backend/config"
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type ansGetTelemetryToken struct {
	Txn            string `fesl:"TXN"`
	TelemetryToken string `fesl:"telemetryToken"`
	Enabled        string `fesl:"enabled"`
	Filters        string `fesl:"filters"`
	Disabled       bool   `fesl:"disabled"`
}

// GetTelemetryToken handles acct.GetTelemetryToken command
func (acct *Account) GetTelemetryToken(event network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansGetTelemetryToken{
			Txn:            acctGetTelemetryToken,
			TelemetryToken: config.General.TelemetryToken,
			Enabled:        "US",
		},
	)
}
