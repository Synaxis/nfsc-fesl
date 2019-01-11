package model

import (
	"github.com/gocraft/dbr"
)

type Queries struct {
	Conn *dbr.Connection
}

type QueriesAdapter interface {
	FindServerByID(sess *dbr.Session, serverID int) (Server, error)
	FindServerBySoldierName(sess *dbr.Session, soldierName string) (Server, error)
	FindServerByCredentials(sess *dbr.Session, accountName string) (Server, error)

	FindPlayerByToken(sess *dbr.Session, token string) (Player, error)
	FindPlayerByID(sess *dbr.Session, playerID int) (Player, error)

	FindHeroesByPlayerID(sess *dbr.Session, playerID int) ([]Hero, error)
	FindHeroByName(sess *dbr.Session, heroName string) (Hero, error)
	FindHeroStats(sess *dbr.Session, heroID int) (HeroStats, error)
	UpdateHeroStats(tx *dbr.Tx, heroID int, pr *HeroStats) error
}
