package network

import (
	"github.com/Synaxis/nfsc-fesl/backend/network/codec"
)

// ClientEvent is the generic struct for events
// by this Client
type ClientEvent struct {
	Name string
	Data interface{}
}

type EventClientCommand struct {
	Client *Client
	// If TLS (theater then we ignore payloadID - it is always 0x0)
	Command *codec.Command
}

func (c *Client) FireClientClose(event ClientEvent) SocketEvent {
	return SocketEvent{
		Name: "client.close",
		Data: EventClientCommand{Client: c},
	}
}

func (c *Client) FireClose() ClientEvent {
	return ClientEvent{Name: "close", Data: c}
}

func (c *Client) FireClientCommand(event ClientEvent) SocketEvent {
	return SocketEvent{
		Name: "client." + event.Name,
		Data: EventClientCommand{
			Client:  c,
			Command: event.Data.(*codec.Command),
		},
	}
}
