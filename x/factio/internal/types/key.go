package types

const (
	// ModuleName is the name of the module
	ModuleName = "factio"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

//nolint
var (
	// Keys for store prefixes
	FactKey       = []byte{0x11} // prefix for Facts
	VotePowerKey  = []byte{0x13} // prefix for the VotePower
	VoteOnFactKey = []byte{0x14} // prefix for the VoteOnFact
)
