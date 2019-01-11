package network

import (
	"fmt"
	"sync"
)

var (
	once  sync.Once
	Lobby *ClientData
)

type ClientData struct {
	Store map[string]*PlayerData
	mu    sync.Mutex
}

// FIXME: Debug only
func InitClientData() {
	once.Do(func() {
		Lobby = NewClientData()
	})
}

// WARNING: It is only for debugging purposes!
func NewClientData() *ClientData {
	return &ClientData{
		Store: map[string]*PlayerData{},
	}
}

func (s *ClientData) Add(lkey string, pd *PlayerData) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.Store[lkey]; ok {
		return fmt.Errorf("PlayerData of lkey %s already exists", lkey)
	}

	s.Store[lkey] = pd
	return nil
}

func (s *ClientData) Get(lkey string) (*PlayerData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pd, ok := s.Store[lkey]
	if !ok {
		return nil, fmt.Errorf("Cannot find a ClientData by the LKEY")
	}
	return pd, nil
}

func (s *ClientData) Delete(lkey string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Store, lkey)
}

type PlayerData struct {
	LobbyKey   string
	ClientName string

	PlayerID int
	HeroID   int
	HeroName string

	ServerID          int
	ServerSoldierName string
	ServerUserName    string
	GameID            int
}

// ServerData contains such keys as:
//
// Custom:
// * "AP"
//
// CGAM:
// * "B-U-alwaysQueue"
// * "B-U-army_balance"
// * "B-U-army_distribution"
// * "B-U-avail_slots_national"
// * "B-U-avail_slots_royal"
// * "B-U-avg_ally_rank"
// * "B-U-avg_axis_rank"
// * "B-U-community_name"
// * "B-U-data_center"
// * "B-U-elo_rank"
// * "B-U-map"
// * "B-U-percent_full"
// * "B-U-server_ip"
// * "B-U-server_port"
// * "B-U-server_state"
// * "B-maxObservers"
// * "B-numObservers"
// * "B-version"
// * "DISABLE-AUTO-DEQUEUE"
// * "HTTYPE"
// * "HXFR"
// * "INT-IP"
// * "INT-PORT"
// * "JOIN"
// * "LID"
// * "MAX-PLAYERS"
// * "NAME"
// * "PORT"
// * "QLEN"
// * "RESERVE-HOST"
// * "RT"
// * "SECRET"
// * "TID"
// * "TYPE"
// * "UGID"
//
// UGAM:
// (The same as in CGAM)
type ServerData map[string]string

func (sd ServerData) Get(key string) string {
	val, ok := sd[key]
	if !ok {
		// logrus.
		// 	WithField("key", key).
		// 	WithField("val", val).
		// 	Debug("ServerData.Get")
	}
	return val
}

func (sd ServerData) Set(key, value string) {
	sd[key] = value
	// logrus.
	// 	WithField("key", key).
	// 	WithField("val", value).
	// 	Debug("ServerData.Set")
}
