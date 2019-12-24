package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeFactDoesNotExist sdk.CodeType = 101
	CodeRepeatedVote     sdk.CodeType = 102
	CodeVotePower        sdk.CodeType = 103
	CodeVoteDoesNotExist sdk.CodeType = 104
)

// ErrFactDoesNotExist is the error for fact not existing
func ErrFactDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeFactDoesNotExist, "Fact does not exist")
}

// ErrRepeatedVote is the error for repeating vote
func ErrRepeatedVote(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeRepeatedVote, "Already Voted on same stance")
}

// ErrVotePower is the error for not enough vote power
func ErrVotePower(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeVotePower, "Not Enough vote power")
}

// ErrVoteDoesNotExist is the error for vote not existing
func ErrVoteDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeVoteDoesNotExist, "Vote does not exist")
}
