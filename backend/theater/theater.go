package theater

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Synaxis/nfsc-fesl/backend/matchmaking"
	"github.com/Synaxis/nfsc-fesl/backend/network"
)

var (
	// thtrGLST = "GLST"
	thtrKICK = "KICK"
	thtrPLVT = "PLVT"
	thtrUBRA = "UBRA"
)

// Theater Handles incoming and outgoing theater communication
type Theater struct {
	mm        *matchmaking.Pool
	socket    *network.Socket
	socketUDP *network.SocketUDP
}

// New creates and starts a new TheaterManager
func New(bind string, mm *matchmaking.Pool) *Theater {
	socket, err := network.NewSocketTCP(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	socketUDP, err := network.NewSocketUDP(bind)
	if err != nil {
		logrus.Fatal(err)
		return nil
	}

	tm := &Theater{
		mm:        mm,
		socket:    socket,
		socketUDP: socketUDP,
	}

	return tm
}

func (tm *Theater) ListenAndServe(ctx context.Context) {
	go tm.Run(ctx)
}

func (tm *Theater) Run(ctx context.Context) {
	for {
		select {
		case event := <-tm.socketUDP.EventChan:
			tm.handleUDP(event)
		case event := <-tm.socket.EventChan:
			tm.handleTLS(event)
		case <-ctx.Done():
			return
		}
	}
}

func (tm *Theater) handleUDP(event network.SocketUDPEvent) {
	switch event.Name {
	case "ECHO":
		tm.ECHO(event)
	default:
		logrus.
			WithFields(logrus.Fields{
				"event": event.Name,
				"data":  event.Data,
			}).
			Warn("theater.UnhandledRequest (UDP)")
	}
}

func (tm *Theater) handleTLS(event network.SocketEvent) {
	ev, ok := event.Data.(network.EventClientCommand)
	if !ok {
		logrus.Error("Logic error: Cannot cast event to network.EventClientCommand")
		return
	}

	// if !ev.Client.IsActive {
	// 	logrus.WithField("command", ev.Command).Warn("Inactive client")
	// 	return
	// }

	switch event.Name {
	case "newClient":
		tm.newClient(ev.Client)
	case "client.command.CONN":
		tm.Connect(ev)
	case "client.command.USER":
		tm.Login(ev)
	case "client.command.LLST":
		tm.GetLobbyList(ev)
		tm.LobbyData(ev)
	case "client.command.GDAT":
		tm.GameData(ev)
	case "client.command.EGAM":
		tm.EnterGame(ev)
	case "client.command.ECNL":
		tm.EnterConnectionLAN(ev)
	case "client.command.CGAM":
		tm.CreateGame(ev)
	case "client.command.UBRA":
		tm.UpdateBracket(ev)
	case "client.command.UGAM":
		tm.UpdateGameData(ev)
	case "client.command.EGRS":
		tm.EnterGameHostResponse(ev)
	case "client.command.PENT":
		tm.PlayerEntered(ev)
	case "client.command.PLVT":
		tm.PlayerExited(ev)
	case "client.command.UPLA":
		tm.UpdatePlayerData(ev)
	case "client.command.PING":
		// TODO: Use metrics in the response and save it for later use.
		return
	default:
		logrus.
			WithFields(logrus.Fields{
				"event":   event.Name,
				"query":   ev.Command.Query,
				"payload": ev.Command.Message,
			}).
			Warn("theater.UnhandledRequest")
	}
}

func (tm *Theater) newClient(client *network.Client) {
	// Start Heartbeat
	client.HeartTicker = time.NewTicker(time.Second * 5)
	go func() {
		for client.IsActive {
			select {
			case <-client.HeartTicker.C:
				if !client.IsActive {
					return
				}
				tm.PING(client)
			}
		}
	}()
}
