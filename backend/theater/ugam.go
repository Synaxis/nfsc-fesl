package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqUGAM struct {
	// TID=14
	TID int `fesl:"TID"`

	// LID=1
	LobbyID int `fesl:"LID"`
	// GID=12
	GameID int `fesl:"GID"`
	// JOIN=O
	JoinMode string `fesl:"JOIN"`
	// MAX-PLAYERS=16
	MaxPlayers int `fesl:"MAX-PLAYERS"`
	// B-maxObservers=0
	MaxObservers int `fesl:"B-maxObservers"`
	// B-numObservers=0
	NumObservers int `fesl:"B-numObservers"`

	// reqUGAMKeys
}

type reqUGAMKeys struct {
	// B-U-army_balance=Axis
	// B-U-army_balance=Balanced
	// B-U-army_distribution="0,0,0,0,0,0,0,0,0,0,0"
	// B-U-army_distribution="1,0,0,1,1,0,0,0,0,0,0"
	// B-U-avail_vips_national=4
	// B-U-avail_vips_royal=4
	// B-U-avg_ally_rank=1000
	// B-U-avg_axis_rank=1000
	// B-U-easyzone=no
	// B-U-elo_rank=1000
	// B-U-lvl_avg=0.000000
	// B-U-lvl_sdv=0.000000
	// B-U-map_name=Village
	// B-U-map_name=seaside_skirmish
	// B-U-percent_full=0
	// B-U-percent_full=6
	// B-U-punkb=0
	// B-U-ranked=yes
	// B-U-server_state=empty
	// B-U-server_state=has_players
	// B-U-servertype=public
	// B-maxObservers=0
	// B-numObservers=0
	// NAME="[iad]A Battlefield Heroes Server(172.28.128.1:18567)"
}

func (tm *Theater) UpdateGameData(event network.EventClientCommand) {
	for key, value := range event.Command.Message {
		switch key {
		case "TID", "LID", "GID":
			// Do nothing
		default:
			event.Client.ServerData.Set(key, value)
		}
	}
}
