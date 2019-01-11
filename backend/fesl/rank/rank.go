package rank

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

const (
	rankGetStats          = "GetStats"
	rankUpdateStats       = "UpdateStats"
	rankGetStatsForOwners = "GetStatsForOwners"
)

// Ranking probably stands for Ranking
type Ranking struct {
}

func (r *Ranking) answer(client *network.Client, pnum uint32, payload interface{}) {
	client.WriteEncode(&codec.Answer{
		Type:         codec.FeslRanking,
		PacketNumber: pnum,
		Payload:      payload,
	})
}
