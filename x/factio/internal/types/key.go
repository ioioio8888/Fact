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
	// Last* values are constant during a block.
	FactKey      = []byte{0x11} // prefix for Facts
	FactCheckKey = []byte{0x12} // prefix for the FactCheck
)
