package pnow

import (
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

type reqStart struct {
	// TXN=Start
	TXN string `fesl:"TXN"`
	// partition.partition=
	Partition statusPartition `fesl:"partition"`
	// debugLevel=off
	DebugLevel string `fesl:"debugLevel"`
	// version=1
	Version int `fesl:"version"`
	// players.[]=1
	Players []reqStartPlayer
}

type reqStartPlayer struct {
	// players.0.ownerId=9
	OwnerID int `fesl:"ownerId"`

	// players.0.ownerType=1
	OwnerType int `fesl:"ownerId"`

	// players.0.props.{sessionType}=listServers
	// players.0.props.{name}=FirstHero
	// players.0.props.{firewallType}=unknown
	// players.0.props.{poolMaxPlayers}=1
	// players.0.props.{poolTimeout}=30
	// players.0.props.{poolTargetPlayers}=0:1
	// players.0.props.{availableServerCount}=1
	// players.0.props.{maxListServersResult}=20
	// players.0.props.{filter-version}=\"1.42.217478.0 \"
	// players.0.props.{filterToGame-version}=version
	// players.0.props.{filter-avail_slots_national}=yes
	// players.0.props.{filterToGame-avail_slots_national}=U-avail_slots_national
	// players.0.props.{filter-data_center}=iad
	// players.0.props.{filterToGame-data_center}=U-data_center
	// players.0.props.{filter-map}=both
	// players.0.props.{filterToGame-map}=U-map
	// players.0.props.{filter-ranked}=yes
	// players.0.props.{filterToGame-ranked}=U-ranked
	// players.0.props.{filter-server_state}=has_players
	// players.0.props.{filterToGame-server_state}=U-server_state
	// players.0.props.{filter-servertype}=public
	// players.0.props.{filterToGame-servertype}=U-servertype
	// players.0.props.{pref-army_balance}=Axis
	// players.0.props.{prefVotingMethod-army_balance}=lottery
	// players.0.props.{fitValues-army_balance}=\"MaxAxis,Axis,Balanced,Allies,MaxAllies\"
	// players.0.props.{fitTable-army_balance}=0;0;0;0;0|-1;0.1;0.5;0.9;1|0;0;0;0;0|1;0.9;0.5;0.1;-1|0;0;0;0;0
	// players.0.props.{fitWeight-army_balance}=200
	// players.0.props.{fitThresholds-army_balance}=0:0
	// players.0.props.{prefToGame-army_balance}=U-army_balance
	// players.0.props.{pref-lvl_avg}=1
	// players.0.props.{aggrPref-lvl_avg}=0
	// players.0.props.{fitScale-lvl_avg}=4
	// players.0.props.{fitWeight-lvl_avg}=200
	// players.0.props.{fitThresholds-lvl_avg}=0:100
	// players.0.props.{prefToGame-lvl_avg}=U-lvl_avg
	// players.0.props.{pref-lvl_sdv}=0
	// players.0.props.{aggrPref-lvl_sdv}=0
	// players.0.props.{fitScale-lvl_sdv}=2
	// players.0.props.{fitWeight-lvl_sdv}=120
	// players.0.props.{fitThresholds-lvl_sdv}=0:0
	// players.0.props.{prefToGame-lvl_sdv}=U-lvl_sdv
	// players.0.props.{pref-percent_full}=80
	// players.0.props.{aggrPref-percent_full}=0
	// players.0.props.{fitScale-percent_full}=30
	// players.0.props.{fitWeight-percent_full}=200
	// players.0.props.{fitThresholds-percent_full}=0:0
	// players.0.props.{prefToGame-percent_full}=U-percent_full
	// players.0.props.{}=44
	Properties map[string]interface{} `fesl:"props"`
}

const (
	debugLevelOff  = "off"
	debugLevelHigh = "high"
	debugLevelMed  = "med"
	debugLevelLow  = "low"
)

type ansStart struct {
	Txn string          `fesl:"TXN"`
	ID  statusPartition `fesl:"id"`
}

// Start handles pnow.Start
func (pnow *PlayNow) Start(event network.EventClientCommand) {
	pnow.answer(
		event.Client,
		event.Command.PayloadID,
		ansStart{
			Txn: pnowStart,
			ID:  statusPartition{1, event.Command.Message["partition.partition"]},
		},
	)
}
