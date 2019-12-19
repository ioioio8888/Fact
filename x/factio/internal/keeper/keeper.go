package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the factio Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Fact metadata struct
func (k Keeper) GetFact(ctx sdk.Context, title string) types.Fact {
	store := ctx.KVStore(k.storeKey)

	key := append(types.FactKey, (title)...)
	bz := store.Get(key)
	if bz == nil {
		//return an empty new fact if fact does not exist
		return types.NewFact()
	}
	var fact types.Fact
	k.cdc.MustUnmarshalBinaryBare(bz, &fact)
	return fact
}

// HasCreator - returns whether or not the title already has an creator
func (k Keeper) HasCreator(ctx sdk.Context, title string) bool {
	return !k.GetFact(ctx, title).Creator.Empty()
}

// Sets the entire fact metadata struct
func (k Keeper) SetFact(ctx sdk.Context, fact types.Fact) {
	if fact.Creator.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	key := append(types.FactKey, (fact.Title)...)
	store.Set(key, k.cdc.MustMarshalBinaryBare(fact))
}

// GetCreator - get the creator of a fact
func (k Keeper) GetCreator(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetFact(ctx, name).Creator
}

// GetPrice - get the price of a fact
func (k Keeper) GetPrice(ctx sdk.Context, title string) sdk.Coins {
	return k.GetFact(ctx, title).Price
}

// Get an iterator over all Fact in which the keys are the names and the values are the Fact
func (k Keeper) GetFactIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FactKey)
}

// function to return delegation
func (k Keeper) GetFactDelegation(ctx sdk.Context, title string, delegator sdk.AccAddress) types.FactDelegation {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FactDelegateKey, (delegator.String() + title)...)
	bz := store.Get(key)
	if bz == nil {
		//return an empty new factdelegation if factdelegation does not exist
		return types.NewDelegation()
	}
	var factDelegation types.FactDelegation
	k.cdc.MustUnmarshalBinaryBare(bz, &factDelegation)
	return factDelegation
}

// function to return if a factdelegation object exist
func (k Keeper) HasFactDelegation(ctx sdk.Context, title string, delegator sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FactDelegateKey, (delegator.String() + title)...)
	bz := store.Get(key)
	if bz == nil {
		return false
	}
	return true
}

// function to set a delegation object
func (k Keeper) SetFactDelegation(ctx sdk.Context, factdelegation types.FactDelegation) {
	if factdelegation.Delegator.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	key := append(types.FactDelegateKey, (factdelegation.Delegator.String() + factdelegation.Title)...)
	store.Set(key, k.cdc.MustMarshalBinaryBare(factdelegation))
}

// Deletes the entire FactDelegation metadata struct
func (k Keeper) DeleteFactDelegation(ctx sdk.Context, title string, delegator sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := append(types.FactDelegateKey, (delegator.String() + title)...)
	store.Delete([]byte(key))
}

// Get an iterator over all Fact in which the keys are the names and the values are the Fact
func (k Keeper) GetFactDelegationIterator(ctx sdk.Context, delegator string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	prefix := append(types.FactDelegateKey, delegator...)
	return sdk.KVStorePrefixIterator(store, prefix)
}

// Get an iterator over all Fact in which the keys are the names and the values are the Fact
func (k Keeper) GetAllFactDelegationIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FactDelegateKey)
}

//return all fact delegations
func (k Keeper) GetAllFactDelegation(ctx sdk.Context) types.FactDelegationList {

	factDelegationList := types.NewDelegationList()
	iterator := k.GetAllFactDelegationIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		var out types.FactDelegation
		k.cdc.UnmarshalBinaryBare(iterator.Value(), &out)
		factDelegationList.Delegations = append(factDelegationList.Delegations, out)
	}

	return factDelegationList
}
