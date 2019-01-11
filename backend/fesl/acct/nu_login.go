package acct

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqNuLogin struct {
	// TXN=NuLogin
	TXN string `fesl:"TXN"`
	// returnEncryptedInfo=0
	ReturnEncryptedInfo int `fesl:"returnEncryptedInfo"`
	// macAddr=$0a0027000000
	MacAddr string `fesl:"macAddr"`
}

type reqNuLoginServer struct {
	reqNuLogin

	AccountName     string `fesl:"nuid"`     // Value specified in +eaAccountName
	AccountPassword string `fesl:"password"` // Value specified in +eaAccountPassword
}

type reqNuLoginClient struct {
	reqNuLogin

	EncryptedInfo string `fesl:"encryptedInfo"` // Value specified in +sessionId
}

type ansNuLogin struct {
	Txn       string `fesl:"TXN"`
	ProfileID int    `fesl:"profileId"`
	UserID    int    `fesl:"userId"`
	NucleusID int    `fesl:"nuid"`
	LobbyKey  string `fesl:"lkey"`
}

type ansNuLoginErr struct {
	Txn     string                `fesl:"TXN"`
	Message string                `fesl:"localizedMessage"`
	Errors  []nuLoginContainerErr `fesl:"errorContainer"`
	Code    int                   `fesl:"errorCode"`
}

type nuLoginContainerErr struct {
	Value      string `fesl:"value"`
	FieldError string `fesl:"fieldError"`
	FieldName  string `fesl:"fieldName"`
}

// NuLogin handles acct.NuLogin command
func (acct *Account) NuLogin(event network.EventClientCommand) {
	uniqueID := uuid.NewV4()
	lkey := uniqueID.String()
	event.Client.PlayerData.LobbyKey = lkey
	err := network.Lobby.Add(lkey, event.Client.PlayerData)
	if err != nil {
		logrus.WithError(err).Warn("Cannot add ClientData in acct.NuLogin")
		return
	}

	switch event.Client.GetClientType() {
	case clientTypeServer:
		acct.serverNuLogin(event)
	default:
		acct.clientNuLogin(event)
	}
}

func (acct *Account) clientNuLogin(event network.EventClientCommand) {
	player := 1
	event.Client.PlayerData.PlayerID = player

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLogin{
			Txn:       acctNuLogin,
			UserID:    event.Client.PlayerData.PlayerID,
			ProfileID: event.Client.PlayerData.PlayerID,
			NucleusID: event.Client.PlayerData.PlayerID,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

func (acct *Account) clientNuLoginNotAuthorized(event *network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginErr{
			Txn:     acctNuLogin,
			Message: `"The user is not entitled to access this game"`,
			Code:    120,
		},
	)
}

// acctNuLoginServer - login command for servers
func (acct *Account) serverNuLogin(event network.EventClientCommand) {
	srv := 1


	event.Client.PlayerData.ServerID = srv
	event.Client.PlayerData.ServerSoldierName = "1234"
	event.Client.PlayerData.ServerUserName = "1234"

	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLogin{
			Txn:       acctNuLogin,
			ProfileID: srv,
			UserID:    srv,
			NucleusID: srv,
			LobbyKey:  event.Client.PlayerData.LobbyKey,
		},
	)
}

func (acct *Account) serverNuLoginNotAuthorized(event *network.EventClientCommand) {
	acct.answer(
		event.Client,
		event.Command.PayloadID,
		ansNuLoginErr{
			Txn:     acctNuLogin,
			Message: `"The password the user specified is incorrect"`,
			Code:    122,
		},
	)
}
