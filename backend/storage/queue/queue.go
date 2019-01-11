package queue

const (
	PlayerConnected = iota
	PlayerLeft
)

func ListenEvents(events <-chan int) {
	// TODO: events is some kind of queue
	for ev := range events {
		switch ev {
		case PlayerConnected:
			// OnPlayerConnected()
		case PlayerLeft:
			// OnPlayerLeft()
		}
	}
}

type EventPlayerConnected struct {
	PlayerID string
}

func OnPlayerConnected(ev EventPlayerConnected) {
	// TODO: Fetch all data from main mysql server about specified PlayerID
	// Then put it into local database
	// When done inform main cluster the task is done,
	// It needs to be done really fast - max 3 sec
	// Inform main cluster the player is tied to specified region
	// And main cluster needs to receive PlayerLeft event to
	// unlock it from connecting to other regions
}

type EventPlayerLeft struct {
	PlayerID string
}

func OnPlayerLeft(ev EventPlayerLeft) {
	// TODO: Put all data to main server and unlock it to connect other region
}
