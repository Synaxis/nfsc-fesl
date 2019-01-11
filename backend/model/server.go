package model

import (
	"github.com/gocraft/dbr"
)

const (
	tableServers = "servers"
)

type Server struct {
	ID int `db:"server_id"`

	APIKey string `db:"api_key"` // Used in heroes-api as a X-Server-Key header

	SoldierName string `db:"soldier_name"` // Used in acct.NuLoginPersona

	AccountUsername string `db:"account_username"` // Used in acct.NuLogin
	AccountPassword string `db:"account_password"` // Used in acct.NuLogin
}

func InsertServer(tx *dbr.Tx, server *Server) error {
	r, err := tx.
		InsertInto(tableServers).
		Columns(
			"soldier_name",
			"api_key",
			"account_username",
			"account_password",
		).
		Record(server).
		Exec()
	if err != nil {
		return err
	}
	i, err := r.LastInsertId()
	if err != nil {
		return err
	}
	server.ID = int(i)
	return nil
}

func (q *Queries) FindServerByID(sess *dbr.Session, serverID int) (server Server, err error) {
	err = sess.
		Select(
			"server_id",
			"soldier_name",
			"account_username",
		).
		From(tableServers).
		Where("server_id = ?", serverID).
		LoadOne(&server)
	return server, err
}

func (q *Queries) FindServerBySoldierName(sess *dbr.Session, soldierName string) (server Server, err error) {
	err = sess.
		Select(
			"server_id",
			"soldier_name",
			"account_username",
		).
		From(tableServers).
		Where("soldier_name = ?", soldierName).
		LoadOne(&server)
	return server, err
}

func (q *Queries) FindServerByCredentials(sess *dbr.Session, accountName string) (server Server, err error) {
	err = sess.
		Select(
			"server_id",
			"soldier_name",
			"account_username",
			"account_password",
		).
		From(tableServers).
		Where("account_username = ?", accountName).
		LoadOne(&server)
	return server, err
}
