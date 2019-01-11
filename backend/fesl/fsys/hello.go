package fsys

import (
	"time"

	"github.com/Synaxis/nfsc-fesl/backend/config"
	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

// reqHello is definition of the fsys.Hello request call
// Packet ID: [192 0 0 1]
type reqHello struct {
	// TXN stands for Taxon, sub-query name of the command.
	// TXN=Hello
	TXN string `fesl:"TXN"`
	// clientString=bfwest-pc
	ClientString string `fesl:"clientString"`
	// sku=125170
	Sku int `fesl:"sku"`
	// locale=en_US
	Locale string `fesl:"locale"`
	// clientPlatform=PC
	ClientPlatform string `fesl:"clientPlatform"`
	// clientVersion="1.42.217478.0 "
	ClientVersion string `fesl:"clientVersion"`
	// SDKVersion=5.0.0.0.0
	SdkVersion string `fesl:"SDKVersion"`
	// protocolVersion=2.0
	ProtocolVersion string `fesl:"protocolVersion"`
	// fragmentSize=8096
	FragmentSize int `fesl:"fragmentSize"`
	// clientType=client-noreg
	ClientType string `fesl:"clientType"`
}

type ansHello struct {
	// TXN stands for Taxon, sub-query name of the command.
	// Should be equal: Hello.
	TXN string `fesl:"TXN"`

	// ConnTTL defines how long connection to FESL backend lasts.
	// Thanks to fsys.MemCheck command the connection can be extended
	// for longer duration.
	ConnTTL int `fesl:"activityTimeoutSecs"`

	// ConnectedAt defines when client has connected to backend.
	// This time is used to synchronize client's time with backend.
	ConnectedAt string `fesl:"curTime"`

	// TheaterIP defines an IP address, where Theater backend is located.
	// Probably: It can be a domain name.
	TheaterIP string `fesl:"theaterIp"`

	// TheaterPort defines a Port number on which Theater backend is listening.
	TheaterPort int `fesl:"theaterPort"`

	Domain           domainPartition `fesl:"domainPartition"`
	AddressRemapping string          `fesl:"addressRemapping,omitempty"`

	MessengerIP   string `fesl:"messengerIp"`
	MessengerPort int    `fesl:"messengerPort"`
}

type domainPartition struct {
	Name    string `fesl:"domain"`
	SubName string `fesl:"subDomain"`
}

// Hello handles fsys.Hello command.
//
// It is first command send by game-client and game-server.
func (fsys *ConnectSystem) Hello(client *network.Client, command *codec.Command) {
	// fm.createClient(event.Client.IpAddr.String(), event.Command.Message)
	client.Type = command.Message["clientType"]

	ans := ansHello{
		TXN:           fsysHello,
		ConnTTL:       int((60 * time.Second).Seconds()),
		ConnectedAt:   time.Now().Format("Jan-02-2006 15:04:05 MST"),
		TheaterIP:     config.General.ThtrAddr,
		MessengerIP:   config.General.TelemetryIP,
		MessengerPort: config.General.TelemetryPort,
	}

	if fsys.ServerMode {
		ans.Domain = domainPartition{"eagames", "bfwest-server"}
		ans.TheaterPort = config.General.ThtrServerPort
	} else {
		ans.Domain = domainPartition{"eagames", "bfwest-dedicated"}
		ans.TheaterPort = config.General.ThtrClientPort
	}

	// ans.ErrorCode = (non zero) -> Hello exit with fsys.Goodbye

	fsys.answer(client, 0xC0000001, ans)
}
