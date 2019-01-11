package fsys

type reqSuicide struct {
	// TXN stands for Taxon, sub-query name of the command.
	// Should be equal: Suicide.
	TXN string `fesl:"TXN"`

	Reason string `fesl:"reason"`
}
