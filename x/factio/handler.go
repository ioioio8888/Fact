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
