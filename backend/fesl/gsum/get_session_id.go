package gsum

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type ansGetSessionID struct {
	Txn string `fesl:"TXN"`
	// Games  []Game  `fesl:"games"`
	// Events []Event `fesl:"events"`
}

// GetSessionID handles gsum.GetSessionID command
func (gsum *GameSummary) GetSessionID(client *network.Client, event *codec.Command) {
	gsum.answer(client, 0, ansGetSessionID{
		Txn: gsumGetSessionID,
	})
}
