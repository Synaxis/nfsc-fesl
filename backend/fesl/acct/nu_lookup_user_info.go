package acct

import (
	"strconv"

	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqNuLookupUserInfo struct {
	// TXN=NuLookupUserInfo
	TXN string `fesl:"TXN"`

	// userInfo.[]=1
	// userInfo.0.userName=FirstHero
	UserInfo []userInfo `fesl:"userInfo"`
}

type ansNuLookupUserInfo struct {
	Txn      string     `fesl:"TXN"`
	UserInfo []userInfo `fesl:"userInfo"`
}

type userInfo struct {
	Namespace    string `fesl:"namespace"`
	XBoxUserID   string `fesl:"xuid,omitempty"` // int
	MasterUserID int    `fesl:"masterUserId"`
	UserID       int    `fesl:"userId"`
	UserName     string `fesl:"userName"`
}

// NuLookupUserInfo handles acct.NuLookupUserInfo command
func (acct *Account) NuLookupUserInfo(event network.EventClientCommand) {
	if event.Client.GetClientType() == clientTypeServer {
		if event.Command.Message["userInfo.0.userName"] == event.Client.PlayerData.ServerUserName {
			acct.serverNuLookupUserInfo(event)
			return
		}
	}

	acct.clientNuLookupUserInfo(event)
}

func (acct *Account) clientNuLookupUserInfo(event network.EventClientCommand) {
	heroes := []userInfo{}
	keys, _ := strconv.Atoi(event.Command.Message["userInfo.[]"])
	for i := 0; i < keys; i++ {

		h := 1

		masterHeroID := h

		heroes = append(heroes, userInfo{
			UserName:     "1234",
			UserID:       h,
			MasterUserID: masterHeroID,
			Namespace:    "MAIN",
			XBoxUserID:   "24",
		})
	}

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLookupUserInfo{Txn: acctNuLookupUserInfo, UserInfo: heroes},
	)
}

func (acct *Account) serverNuLookupUserInfo(event network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLookupUserInfo{
			Txn: acctNuLookupUserInfo,
			UserInfo: []userInfo{
				{
					Namespace:    "MAIN",
					XBoxUserID:   "24",
					MasterUserID: event.Client.PlayerData.ServerID,
					UserID:       event.Client.PlayerData.ServerID,
					UserName:     event.Client.PlayerData.ServerUserName,
				},
			},
		},
	)
}
