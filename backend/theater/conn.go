package theater

import (
	"time"

	"github.com/Synaxis/nfsc-fesl/backend/network"
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

type reqCONN struct {
	// TID=1
	TID int `fesl:"TID"`

	// LOCALE=en_US
	Locale string `fesl:"LOCALE"`
	// PLAT=PC"
	Platform string `fesl:"PLAT"`
	// PROD=bfwest-pc
	Prod string `fesl:"PROD"`
	// PROT=2
	Protocol int `fesl:"PROT"`
	// SDKVERSION=5.0.0.0.0
	SdkVersion string `fesl:"SDKVERSION"`
	// VERS="1.42.217478.0 "
	Version string `fesl:"VERS"`
}

type ansCONN struct {
	TID         string `fesl:"TID"`
	ConnectedAt int64  `fesl:"TIME"`

	// ConnTTL defines how long connection to Theater backend lasts.
	// Thanks to PING command the connection can be extended
	// for longer duration.
	ConnTTL  int    `fesl:"activityTimeoutSecs"`
	Protocol string `fesl:"PROT"`
}

func (tm *Theater) Connect(event network.EventClientCommand) {
	event.Client.WriteEncode(&codec.Answer{
		Type: codec.ThtrConnect,
		Payload: ansCONN{
			TID:         event.Command.Message["TID"],
			ConnectedAt: time.Now().UTC().Unix(),
			ConnTTL:     int((60 * time.Second).Seconds()),
			Protocol:    event.Command.Message["PROT"],
		},
	})
}
