package theater

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqUPLA struct {
	// TID=12
	TID int `fesl:"TID"`
	// GID=3
	GameID int `fesl:"GID"`
	// LID=1
	LobbyID string `fesl:"LID"`
	// PID=6
	PlayerID int `fesl:"PID"`
	// HMO=6
	HostOwnerID int `fesl:"HMO"`

	reqUPLAKeys
}

type reqUPLAKeys struct {
	PlayerElo   *string `fesl:"P-elo"`
	PlayerKills *string `fesl:"P-kills"`
	PlayerKit   *string `fesl:"P-kit"`
	PlayerLevel *string `fesl:"P-level"`
	// P-ping=24
	PlayerPing  *int    `fesl:"P-ping"`
	PlayerScore *string `fesl:"P-score"`
	PlayerTeam  *string `fesl:"P-team"`
	// P-time="1 min 10 sec "
	PlayerPlayedTime *string `fesl:"P-time"`
	PlayerClientID   *string `fesl:"P-cid"`
	PlayerDataCenter *string `fesl:"P-dc"`
	PlayerIP         *string `fesl:"P-ip"`
}

// type ansUPLA struct {
// LobbyID string `fesl:"LID"`
// 	TID int `fesl:"TID"`

// 	PlayerID int `fesl:"PID"`
// 	Name string `fesl:"NAME"`
// 	UserID string `fesl:"UID"` // optional

// 	reqUPLAKeys
// }

func (tm *Theater) UpdatePlayerData(event network.EventClientCommand) {
	num, _ := strconv.Atoi(event.Client.ServerData.Get("AP"))
	num++
	event.Client.ServerData.Set("AP", strconv.Itoa(num))

	logrus.
		WithField("payload", event.Command.Message).
		WithField("playerID", event.Command.Message.Get("PID")).
		Println("UPLA")
}
