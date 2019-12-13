package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgCreateFact defines a Create Fact Message
type MsgCreateFact struct {
	Title       string         `json:"title"`
	Bid         sdk.Coins      `json:"bid"`
	Creator     sdk.AccAddress `json:"creator"`
	Time        int64          `json:"time"`
	Place       string         `json:"place"`
	Description string         `json:"description"`
}

// NewMsgCreateFact is a constructor function for MsgCreateFact
func NewMsgCreateFact(title string, creator sdk.AccAddress, bid sdk.Coins, time int64, place string, description string) MsgCreateFact {
	return MsgCreateFact{
		Title:       title,
		Bid:         bid,
		Creator:     creator,
		Time:        time,
		Place:       place,
		Description: description,
	}
}

// Route should return the name of the module
func (msg MsgCreateFact) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateFact) Type() string { return "create_fact" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateFact) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(msg.Creator.String())
	}
	if len(msg.Title) == 0 {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}
	if len(msg.Title) >= 60 {
		return sdk.ErrUnknownRequest("title cannot be more than 60 words")
	}
	if len(msg.Description) == 0 {
		return sdk.ErrUnknownRequest("description cannot be empty")
	}
	if len(msg.Description) >= 280 {
		return sdk.ErrUnknownRequest("description cannot be more than 280 words")
	}
	if msg.Time == 0 {
		return sdk.ErrUnknownRequest("Time cannot be empty")
	}
	if len(msg.Place) == 0 {
		return sdk.ErrUnknownRequest("Place cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateFact) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateFact) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// MsgEditFact defines a Edit Fact Message
type MsgEditFact struct {
	Title       string         `json:"title"`
	Editor      sdk.AccAddress `json:"editor"`
	Time        int64          `json:"time"`
	Place       string         `json:"place"`
	Description string         `json:"description"`
}

// NewMsgEditFact is a constructor function for MsgEditFact
func NewMsgEditFact(title string, editor sdk.AccAddress, bid sdk.Coins, time int64, place string, description string) MsgEditFact {
	return MsgEditFact{
		Title:       title,
		Editor:      editor,
		Time:        time,
		Place:       place,
		Description: description,
	}
}

// Route should return the name of the module
func (msg MsgEditFact) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditFact) Type() string { return "edit_fact" }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditFact) ValidateBasic() sdk.Error {
	if msg.Editor.Empty() {
		return sdk.ErrInvalidAddress(msg.Editor.String())
	}
	if len(msg.Title) == 0 {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}
	if len(msg.Title) >= 60 {
		return sdk.ErrUnknownRequest("title cannot be more than 60 words")
	}
	if len(msg.Description) == 0 {
		return sdk.ErrUnknownRequest("description cannot be empty")
	}
	if len(msg.Description) >= 280 {
		return sdk.ErrUnknownRequest("description cannot be more than 280 words")
	}
	if msg.Time == 0 {
		return sdk.ErrUnknownRequest("Time cannot be empty")
	}
	if len(msg.Place) == 0 {
		return sdk.ErrUnknownRequest("Place cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditFact) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditFact) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Editor}
}

// MsgEditFact defines a Edit Fact Message
type MsgDelegateFact struct {
	Title       string         `json:"title"`
	Editor      sdk.AccAddress `json:"editor"`
	Time        int64          `json:"time"`
	Place       string         `json:"place"`
	Description string         `json:"description"`
}

// NewMsgDelegateFact is a constructor function for MsgDelegateFact
func NewMsgDelegateFact(title string, editor sdk.AccAddress, bid sdk.Coins, time int64, place string, description string) MsgEditFact {
	return MsgEditFact{
		Title:  title,
		Editor: editor,
		// Amount: amount,
	}
}

// Route should return the name of the module
func (msg MsgDelegateFact) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDelegateFact) Type() string { return "delegate_fact" }

// ValidateBasic runs stateless checks on the message
func (msg MsgDelegateFact) ValidateBasic() sdk.Error {
	if msg.Editor.Empty() {
		return sdk.ErrInvalidAddress(msg.Editor.String())
	}
	if len(msg.Title) == 0 {
		return sdk.ErrUnknownRequest("Title cannot be empty")
	}
	if len(msg.Title) >= 60 {
		return sdk.ErrUnknownRequest("title cannot be more than 60 words")
	}
	if len(msg.Description) == 0 {
		return sdk.ErrUnknownRequest("description cannot be empty")
	}
	if len(msg.Description) >= 280 {
		return sdk.ErrUnknownRequest("description cannot be more than 280 words")
	}
	if msg.Time == 0 {
		return sdk.ErrUnknownRequest("Time cannot be empty")
	}
	if len(msg.Place) == 0 {
		return sdk.ErrUnknownRequest("Place cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDelegateFact) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDelegateFact) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Editor}
}
