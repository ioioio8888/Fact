package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

// function to return if a Vote On Fact object exist
func (k Keeper) HasVoteOnFact(ctx sdk.Context, title string, voter sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VoteOnFactKey, (voter.String() + title)...)
	bz := store.Get(key)
	if bz == nil {
		return false
	}
	return true
}

// function to return vote on fact
func (k Keeper) GetVoteOnFact(ctx sdk.Context, title string, voter sdk.AccAddress) types.VoteOnFact {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VoteOnFactKey, (voter.String() + title)...)
	bz := store.Get(key)
	if bz == nil {
		//return an empty new factdelegation if factdelegation does not exist
		return types.NewVoteOnFact()
	}
	var voteOnFact types.VoteOnFact
	k.cdc.MustUnmarshalBinaryBare(bz, &voteOnFact)
	return voteOnFact
}

// function to set a vote on fact object
func (k Keeper) SetVoteOnFact(ctx sdk.Context, voteOnFact types.VoteOnFact) {
	if voteOnFact.Voter.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	key := append(types.VoteOnFactKey, (voteOnFact.Voter.String() + voteOnFact.Title)...)
	store.Set(key, k.cdc.MustMarshalBinaryBare(voteOnFact))
}

// Deletes the entire VoteOnFact metadata struct
func (k Keeper) DeleteVoteOnFact(ctx sdk.Context, title string, voter sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.VoteOnFactKey, (voter.String() + title)...)
	store.Delete([]byte(key))
}

// Get an iterator over all vote on fact which the key starts with the voter
func (k Keeper) GetVoteOnFactIterator(ctx sdk.Context, voter string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.VoteOnFactKey, voter...)
	return sdk.KVStorePrefixIterator(store, prefix)
}
