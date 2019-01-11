package theater

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqEGEG struct {
	reqEGAM
}

type ansEGEG struct {
	TID           string `fesl:"TID"`
	Platform      string `fesl:"PL"`
	Ticket        string `fesl:"TICKET"`
	PlayerID      int    `fesl:"PID"`
	IP            string `fesl:"I"`
	Huid          string `fesl:"HUID"`
	Port          string `fesl:"P"`
	HostUserID    int    `fesl:"HUID"`
	EncryptionKey string `fesl:"EKEY"`
	// Alternatively to EKEY it is possible to use NOENCYRPTIONKEY
	NoEcryptionKey string `fesl:"NOENCYRPTIONKEY,omitempty"`
	IntIP          string `fesl:"INT-IP"`
	IntPort        string `fesl:"INT-PORT"`
	Secret         string `fesl:"SECRET,omitempty"`
	// Alternatively to SECRET it is possible to use NOSECRET
	NoSecret string `fesl:"NOSECRET,omitempty"`
	Ugid     string `fesl:"UGID,omitempty"`
	Xuid     string `fesl:"XUID"`
    RXuid    string `fesl:"R-XUID"`
	// Alternatively to UGID it is possible to use NOGUID
	NoGUID  string `fesl:"NOGUID,omitempty"`
	LobbyID string `fesl:"LID"`
	GameID  int    `fesl:"GID"`
}

// EGEG is sent Client to receive last confirmation before joining game
func (tm *Theater) EGEG(event *network.EventClientCommand, gameServer *network.Client, gr GameRequest) {
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrEnterGameEntitleGame,
		Payload: ansEGEG{
			TID:           event.Command.Message["TID"],
			Platform:      "PC",
			Huid:    		"1",
			Xuid:         "24",
			RXuid:        "24",
			Ticket:        gr.Ticket,
			PlayerID:      gr.PlayerID,
			IP:            gameServer.ServerData.Get("IP"),
			Port:          gameServer.ServerData.Get("PORT"),
			HostUserID:    gameServer.PlayerData.ServerID,
			EncryptionKey: "O65zZ2D2A58mNrZw1hmuJw%3d%3d",
			// NoEcryptionKey: "1",
			IntIP:   gameServer.ServerData.Get("INT-IP"),
			IntPort: gameServer.ServerData.Get("INT-PORT"),
			Secret:  "MargeSimpson",
			Ugid:    gameServer.ServerData.Get("UGID"),
			LobbyID: gr.LobbyID,
			GameID:  gr.GameID,
		},
	})
}
