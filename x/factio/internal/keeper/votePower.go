package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

// function to get a vote power object
func (k Keeper) GetVotePower(ctx sdk.Context, voter sdk.AccAddress) types.VotePower {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VotePowerKey, voter.String()...)
	bz := store.Get(key)
	if bz == nil {
		//return an empty new factdelegation if factdelegation does not exist
		return types.NewVotePower()
	}
	var votePower types.VotePower
	k.cdc.MustUnmarshalBinaryBare(bz, &votePower)
	return votePower
}

// function to set a votepower object
func (k Keeper) SetVotePower(ctx sdk.Context, votePower types.VotePower) {
	if votePower.Voter.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	key := append(types.VotePowerKey, votePower.Voter.String()...)
	store.Set(key, k.cdc.MustMarshalBinaryBare(votePower))
}

// function to return if a votepower object exist
func (k Keeper) HasVotePower(ctx sdk.Context, voter sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VotePowerKey, voter.String()...)
	bz := store.Get(key)
	if bz == nil {
		return false
	}
	return true
}

//function to create a new votepower object to an address which has never voted before
func (k Keeper) CreateVotePower(ctx sdk.Context, voter sdk.AccAddress) {
	NewVotePower := types.NewVotePower()
	NewVotePower.Voter = voter
	var amount sdk.Dec
	hasFactCoin := false
	for _, coin := range k.CoinKeeper.GetCoins(ctx, voter) {
		if coin.Denom == "factcoin" {
			hasFactCoin = true
			amount = coin.Amount.ToDec()
		}
	}
	if hasFactCoin {
		NewVotePower.Power = amount
	} else {
		NewVotePower.Power = sdk.NewDec(0)
	}
	k.SetVotePower(ctx, NewVotePower)
}

// Get an iterator over all votepower
func (k Keeper) GetAllVotePowerIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.VotePowerKey)
}

//return all Vote power
func (k Keeper) GetAllVotePower(ctx sdk.Context) types.VotePowerList {

	votePowerList := types.NewVotePowerList()
	iterator := k.GetAllVotePowerIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		var out types.VotePower
		k.cdc.UnmarshalBinaryBare(iterator.Value(), &out)
		votePowerList.VotePowers = append(votePowerList.VotePowers, out)
	}

	return votePowerList
}

//function to restore the vote power on every votepower object
func (k Keeper) RestoreVotePower(ctx sdk.Context) {
	votePowerList := k.GetAllVotePower(ctx)
	for _, votePower := range votePowerList.VotePowers {
		var amount sdk.Dec
		hasFactCoin := false
		for _, coin := range k.CoinKeeper.GetCoins(ctx, votePower.Voter) {
			if coin.Denom == "factcoin" {
				hasFactCoin = true
				amount = coin.Amount.ToDec()
			}
		}
		if hasFactCoin {
			votePower.Power = amount
		} else {
			votePower.Power = sdk.NewDec(0)
		}
		k.SetVotePower(ctx, votePower)
	}
}
