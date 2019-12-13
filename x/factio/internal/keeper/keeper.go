package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/sdk-tutorials/factio/x/factio/internal/types"
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

// Gets the entire Fact metadata struct for a name
func (k Keeper) GetFact(ctx sdk.Context, title string) types.Fact {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(title))
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

// Sets the entire Whois metadata struct for a name
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
