package theater

import (
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqUSER struct {
	// TID=2
	TID string `fesl:"TID"`
	// LKEY=3c63a203-80d5-462a-9112-414345d40376
	LobbyKey string `fesl:"LKEY"`
	// CID=
	ClientID string `fesl:"CID"`
	// MAC=$0a0027000000
	MACAddr string `fesl:"MAC"`
	// SKU=125170
	SKU string `fesl:"SKU"`
	// NAME=
	NAME string `fesl:"NAME"`
}

type ansUSER struct {
	TID      string `fesl:"TID"`
	Name     string `fesl:"NAME"`
	ClientID string `fesl:"CID,omitempty"`
}

func (tm *Theater) Login(event network.EventClientCommand) {
	cd, err := network.Lobby.Get(event.Command.Message.Get("LKEY"))
	if err != nil {
		logrus.WithError(err).Warn("Cannot find client data in theater.USER")
		return
	}
	// FIXME Also server data
	event.Client.PlayerData = cd

	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrUser,
		Payload: ansUSER{
			ClientID: cd.HeroName,
			TID:      event.Command.Message["TID"],
			// Name: event.Client.PlayerData.ClientName,
			Name: cd.HeroName,
		},
	})
}
