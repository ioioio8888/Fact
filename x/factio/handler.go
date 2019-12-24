package factio

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

// NewHandler returns a handler for "factio" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreateFact:
			return handleMsgCreateFact(ctx, keeper, msg)
		case types.MsgEditFact:
			return handleMsgEditFact(ctx, keeper, msg)
		case types.MsgVoteFact:
			return handleMsgVoteFact(ctx, keeper, msg)
		case types.MsgUnVoteFact:
			return handleMsgUnVoteFact(ctx, keeper, msg)
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

//handlge a msg to vote on fact
func handleMsgVoteFact(ctx sdk.Context, keeper Keeper, msg types.MsgVoteFact) sdk.Result {
	fact := keeper.GetFact(ctx, msg.Title)
	if fact.Creator.Empty() {
		return types.ErrFactDoesNotExist("Fact does not exist").Result()
	}
	//check if the address is voted on same side before
	if keeper.HasVoteOnFact(ctx, msg.Title, msg.Voter) {
		if keeper.GetVoteOnFact(ctx, msg.Title, msg.Voter).Upvote == msg.UpVote {
			return types.ErrRepeatedVote("Repeated Vote").Result()
		}
	}
	if !keeper.HasVotePower(ctx, msg.Voter) {
		keeper.CreateVotePower(ctx, msg.Voter)
	}
	votepower := keeper.GetVotePower(ctx, msg.Voter)
	//check if the address has enough vote power
	if sdk.NewDec(1).GT(votepower.Power) {
		return types.ErrVotePower("Not enough vote power").Result()
	}
	//deduct 1 power from the address's votepower
	votepower.Power = votepower.Power.Sub(sdk.NewDec(1))
	keeper.SetVotePower(ctx, votepower)

	//set vote on fact
	voteOnFact := keeper.GetVoteOnFact(ctx, msg.Title, msg.Voter)
	voteOnFact.Voter = msg.Voter
	voteOnFact.Title = msg.Title
	voteOnFact.Upvote = msg.UpVote
	keeper.SetVoteOnFact(ctx, voteOnFact)

	//set the fact
	if msg.UpVote {
		//get the fact and remove the vote that if the address has voted on different stance
		fact := keeper.GetFact(ctx, msg.Title)
		if contains(fact.Downvoters, msg.Voter) {
			fact.Downvoters = RemoveIndex(fact.Downvoters, msg.Voter)
		}
		//set the upvoters in the list
		fact.Upvoters = append(fact.Upvoters, msg.Voter)
		keeper.SetFact(ctx, fact)
	} else {
		//get the fact and remove the vote that if the address hason voted on different stance
		fact := keeper.GetFact(ctx, msg.Title)
		if contains(fact.Upvoters, msg.Voter) {
			fact.Upvoters = RemoveIndex(fact.Upvoters, msg.Voter)
		}
		//set the downvoters in the list
		fact.Downvoters = append(fact.Downvoters, msg.Voter)
		keeper.SetFact(ctx, fact)
	}
	return sdk.Result{}
}

func contains(s []sdk.AccAddress, e sdk.AccAddress) bool {
	for _, a := range s {
		if a.Equals(e) {
			return true
		}
	}
	return false
}

//handlge a msg to unvote fact
func handleMsgUnVoteFact(ctx sdk.Context, keeper Keeper, msg types.MsgUnVoteFact) sdk.Result {
	fact := keeper.GetFact(ctx, msg.Title)
	if fact.Creator.Empty() {
		return types.ErrFactDoesNotExist("Fact does not exist").Result()
	}
	//check if the address is voted on this fact before, if no, return an error
	if !keeper.HasVoteOnFact(ctx, msg.Title, msg.Voter) {
		return types.ErrVoteDoesNotExist("Vote does not exist").Result()
	}

	vote := keeper.GetVoteOnFact(ctx, msg.Title, msg.Voter)
	//Remove the address stored in upvoters/downvoters
	if vote.Upvote {
		fact.Upvoters = RemoveIndex(fact.Upvoters, msg.Voter)
	} else if !vote.Upvote {
		fact.Downvoters = RemoveIndex(fact.Downvoters, msg.Voter)
	}
	keeper.SetFact(ctx, fact)
	//delete the vote on Fact obeject
	keeper.DeleteVoteOnFact(ctx, msg.Title, msg.Voter)
	return sdk.Result{}
}
