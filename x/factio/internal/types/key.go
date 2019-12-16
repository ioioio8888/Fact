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
	FactKey         = []byte{0x11} // prefix for Facts
	FactDelegateKey = []byte{0x12} // prefix for the FactCheck
)
