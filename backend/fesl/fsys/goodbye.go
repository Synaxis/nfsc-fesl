package fsys

type reqGoodbye struct {
	// TXN=Goodbye
	TXN string `fesl:"TXN"`
	// reason=GOODBYE_CLIENT_NORMAL
	Reason string `fesl:"reason"`
	// message=n/a
	Message string `fesl:"message"`
}
