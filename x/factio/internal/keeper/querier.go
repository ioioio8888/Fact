package keeper

import (
	"bytes"

	"github.com/ioioio8888/factio/x/factio/internal/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the factio Querier
const (
	QueryFact         = "getFact"
	QueryFactList     = "getFactList"
	QueryAccountCoins = "getAccCoins"
	QueryVotePower    = "getVotePower"
	QueryVotedList    = "getVoteList"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryFact:
			return queryFact(ctx, path[1:], req, keeper)
		case QueryFactList:
			return queryFactList(ctx, req, keeper)
		case QueryAccountCoins:
			return queryAccountCoins(ctx, path[1:], req, keeper)
		case QueryVotePower:
			return queryVotePower(ctx, path[1:], req, keeper)
		case QueryVotedList:
			return queryVotedList(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown factio query endpoint")
		}
	}
}

// nolint: unparam
func queryFact(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	if !keeper.HasCreator(ctx, path[0]) {
		panic("Fact does not created yet")
	}

	fact := keeper.GetFact(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, fact)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// return a list of title of all facts
func queryFactList(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var factList types.QueryResFactList
	iterator := keeper.GetFactIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		title := bytes.Trim(iterator.Key(), string(types.FactKey))
		factList = append(factList, string(title))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, factList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// return an account's coins
func queryAccountCoins(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var coins sdk.Coins
	address, aerr := sdk.AccAddressFromBech32(path[0])
	if aerr != nil {
		panic("address format is not correct")
	}
	coins = keeper.CoinKeeper.GetCoins(ctx, address)
	res, err := codec.MarshalJSONIndent(keeper.cdc, coins)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// return an account's vote power
func queryVotePower(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var votePower types.VotePower
	address, aerr := sdk.AccAddressFromBech32(path[0])
	if aerr != nil {
		panic("address format is not correct")
	}
	votePower = keeper.GetVotePower(ctx, address)
	res, err := codec.MarshalJSONIndent(keeper.cdc, votePower)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// return a list of account's voted fact
func queryVotedList(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var votedList types.QueryResVotedList
	iterator := keeper.GetVoteOnFactIterator(ctx, path[0])

	for ; iterator.Valid(); iterator.Next() {
		var out types.VoteOnFact
		keeper.cdc.UnmarshalBinaryBare(iterator.Value(), &out)
		votedList = append(votedList, out)
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, votedList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
