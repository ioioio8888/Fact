package factio

import (
	"fmt"

	"github.com/cosmos/sdk-tutorials/factio/x/factio/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "factio" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreateFact:
			return handleMsgCreateFact(ctx, keeper, msg)
		case types.MsgEditFact:
			return handleMsgEditFact(ctx, keeper, msg)
		case types.MsgDelegateFact:
			return handleMsgDelegateFact(ctx, keeper, msg)
		case types.MsgUnDelegateFact:
			return handleMsgUnDelegateFact(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized factio Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to create Fact
func handleMsgCreateFact(ctx sdk.Context, keeper Keeper, msg types.MsgCreateFact) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Title).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasCreator(ctx, msg.Title) {
		return sdk.ErrInternal("Same Title exists!").Result()
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Creator, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}

	fact := keeper.GetFact(ctx, msg.Title)
	fact.Title = msg.Title
	fact.Time = msg.Time
	fact.Place = msg.Place
	fact.Description = msg.Description
	fact.Creator = msg.Creator
	keeper.SetFact(ctx, fact)
	return sdk.Result{}
}

// Handle a message to edit Fact
func handleMsgEditFact(ctx sdk.Context, keeper Keeper, msg types.MsgEditFact) sdk.Result {

	if !msg.Editor.Equals(keeper.GetCreator(ctx, msg.Title)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Editor is not the Creator").Result() // If not, throw an error
	}

	fact := keeper.GetFact(ctx, msg.Title)
	fact.Title = msg.Title
	fact.Time = msg.Time
	fact.Place = msg.Place
	fact.Description = msg.Description
	keeper.SetFact(ctx, fact)
	return sdk.Result{}
}

// Handle a message to delegate Fact
func handleMsgDelegateFact(ctx sdk.Context, keeper Keeper, msg types.MsgDelegateFact) sdk.Result {

	if !keeper.HasFactDelegation(ctx, msg.Title, msg.Delegator) {
		return sdk.ErrInvalidAddress("This address has already delegated on this fact").Result()
	}
	fcoin, _ := sdk.ParseCoins("1factcoin")
	_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Delegator, fcoin) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}
	dfcoin, _ := sdk.ParseCoins("1delegatedfactcoin")
	_, err = keeper.CoinKeeper.AddCoins(ctx, msg.Delegator, dfcoin) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInvalidCoins("Error adding coins").Result()
	}
	//get the fact delegation and set the fact delegation object
	factdelegation := keeper.GetFactDelegation(ctx, msg.Title, msg.Delegator)
	factdelegation.Delegator = msg.Delegator
	factdelegation.Title = msg.Title
	factdelegation.Shares = factdelegation.Shares.Add(fcoin)
	keeper.SetFactDelegation(ctx, factdelegation)

	//get the fact and add the delegator to the list in fact
	fact := keeper.GetFact(ctx, msg.Title)
	fact.Delegators = append(fact.Delegators, msg.Delegator)
	keeper.SetFact(ctx, fact)

	return sdk.Result{}
}

//find the index of the address from the list
func indexOf(element sdk.AccAddress, data []sdk.AccAddress) int {
	for k, v := range data {
		if element.Equals(v) {
			return k
		}
	}
	return -1 //not found.
}

//remove the address from the list of address
func RemoveIndex(s []sdk.AccAddress, staker sdk.AccAddress) []sdk.AccAddress {
	index := indexOf(staker, s)
	return append(s[:index], s[index+1:]...)
}

// Handle a message to delegate Fact
func handleMsgUnDelegateFact(ctx sdk.Context, keeper Keeper, msg types.MsgUnDelegateFact) sdk.Result {

	if keeper.HasFactDelegation(ctx, msg.Title, msg.Delegator) {
		return sdk.ErrInvalidAddress("This address hasn't delegated on this fact").Result()
	}
	dfcoin, _ := sdk.ParseCoins("1delegatedfactcoin")
	_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Delegator, dfcoin) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInvalidCoins("Error adding coins").Result()
	}
	fcoin, _ := sdk.ParseCoins("1factcoin")
	_, err = keeper.CoinKeeper.AddCoins(ctx, msg.Delegator, fcoin) // If so, deduct the Bid amount from the sender
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}

	keeper.DeleteFactDelegation(ctx, msg.Title, msg.Delegator)

	//get the fact and remove the delegator to the list in fact
	fact := keeper.GetFact(ctx, msg.Title)
	fact.Delegators = RemoveIndex(fact.Delegators, msg.Delegator)
	keeper.SetFact(ctx, fact)

	return sdk.Result{}
}
