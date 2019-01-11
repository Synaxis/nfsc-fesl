package theater

import (
	"net"
	"strconv"

	"github.com/Synaxis/nfsc-fesl/backend/matchmaking"
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

// ->N CGAM 0x40000000
type reqCGAM struct {
	// TID=3
	Tid int `fesl:"TID"`
	// LID=-1
	LobbyID int `fesl:"LID"`
	// RESERVE-HOST=0
	ReserveHost int `fesl:"RESERVE-HOST"`
	// NAME="[iad]A Battlefield Heroes Server(172.28.128.1:18567)"
	Name string `fesl:"NAME"`
	// PORT=18567
	Port int `fesl:"PORT"`
	// HTTYPE=A
	Httype string `fesl:"HTTYPE"`
	// TYPE=G
	Type string `fesl:"TYPE"`
	// QLEN=16
	Qlen int `fesl:"QLEN"`
	// DISABLE-AUTO-DEQUEUE=1
	DisableAutoDequeue int `fesl:"DISABLE-AUTO-DEQUEUE"`
	// HXFR=0
	Hxfr int `fesl:"HXFR"`
	// INT-PORT=18567
	IntPort int `fesl:"INT-PORT"`
	// INT-IP=192.168.1.102
	IntIP string `fesl:"INT-IP"`
	// MAX-PLAYERS=16
	MaxPlayers int `fesl:"MAX-PLAYERS"`
	// B-maxObservers=0
	BMaxObservers int `fesl:"B-maxObservers"`
	// B-numObservers=0
	BNumObservers int `fesl:"B-numObservers"`
	// UGID=GUID-Server
	Ugid string `fesl:"UGID"` /// Value passed in +guid
	// SECRET=Test-Server
	Secret string `fesl:"SECRET"` // Value passed in +secret
	// B-U-alwaysQueue=1
	BUAlwaysQueue int `fesl:"B-U-alwaysQueue"`
	// B-U-army_balance=Balanced
	BUArmyBalance string `fesl:"B-U-army_balance"`
	// B-U-army_distribution="0,0,0, 0,0,0,0, 0,0,0,0"
	BUArmyDistribution string `fesl:"B-U-army_distribution"`
	// B-U-avail_slots_national=yes
	BUAvailSlotsNational string `fesl:"B-U-avail_slots_national"`
	// B-U-avail_slots_royal=yes
	BUAvailSlotsRoyal string `fesl:"B-U-avail_slots_royal"`
	// B-U-avg_ally_rank=1000.0000
	BUAvgAllyRank string `fesl:"B-U-avg_ally_rank"`
	// B-U-avg_axis_rank=1000.0000
	BUAvgAxisRank string `fesl:"B-U-avg_axis_rank"`
	// B-U-community_name="Heroes SV"
	BUCommunityName string `fesl:"B-U-community_name"`
	// B-U-data_center=iad
	BUDataCenter string `fesl:"B-U-data_center"`
	// B-U-elo_rank=1000.0000
	BUEloRank string `fesl:"B-U-elo_rank"`
	// B-U-map=no_vehicles
	BUMap string `fesl:"B-U-map"`
	// B-U-percent_full=0
	BUPercentFull int `fesl:"B-U-percent_full"`
	// B-U-server_ip=172.28.128.1
	BUServerIP string `fesl:"B-U-server_ip"`
	// B-U-server_port=18567
	BUServerPort int `fesl:"B-U-server_port"`
	// B-U-server_state=empty
	BUServerState string `fesl:"B-U-server_state"`
	// B-version=1.46.222034.0
	BVersion string `fesl:"B-version"`
	// JOIN=O
	Join string `fesl:"JOIN"`
	// RT=
	Rt string `fesl:"RT"`
}

const (
	joinCGAMOpen   = "O"
	joinCGAMClosed = "C"
	joinCGAMWait   = "W"
)

type ansCGAM struct {
	TID           string `fesl:"TID"`
	LobbyID       int    `fesl:"LID"`
	MaxPlayers    string `fesl:"MAX-PLAYERS"`
	UGID          string `fesl:"UGID"`
	Secret        string `fesl:"SECRET"`
	JOIN          string `fesl:"JOIN"`
	GameID        int    `fesl:"GID"`
	EncryptionKey string `fesl:"EKEY"`

	// TOOD: It is not used in TheaterCreateGameResult
	J string `fesl:"J"` // wtf ?
}

func (tm *Theater) CreateGame(event network.EventClientCommand) {
	game := &matchmaking.Game{
		GameServer: event.Client,

		// TODO: Use value defined in "MAX-PLAYERS"
		PlayerSlots: 16,
	}
	tm.mm.AddGame(game)

	addr, _ := event.Client.IpAddr.(*net.TCPAddr)
	event.Client.ServerData.Set("LID", "1")
	event.Client.ServerData.Set("GID", strconv.Itoa(game.ID))
	event.Client.ServerData.Set("IP", addr.IP.String())
	event.Client.ServerData.Set("AP", "0")
	event.Client.ServerData.Set("QUEUE-LENGTH", "0")
	for key, value := range event.Command.Message {
		event.Client.ServerData.Set(key, value)
	}

	event.Client.PlayerData.GameID = game.ID

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrCreateGame,
		Payload: ansCGAM{
			TID:           event.Command.Message["TID"],
			LobbyID:       game.LobbyID,
			UGID:          event.Command.Message["UGID"],
			MaxPlayers:    event.Command.Message["MAX-PLAYERS"],
			EncryptionKey: `O65zZ2D2A58mNrZw1hmuJw%3d%3d`,
			// Secret:     `2587913`,
			Secret: event.Command.Message["SECRET"],
			JOIN:   event.Command.Message["JOIN"],
			J:      event.Command.Message["JOIN"],
			GameID: game.ID,
		},
	})
}
