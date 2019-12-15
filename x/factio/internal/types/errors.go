package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeFactDoesNotExist sdk.CodeType = 101
)

// ErrFactDoesNotExist is the error for fact not existing
func ErrFactDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeFactDoesNotExist, "Fact does not exist")
}
