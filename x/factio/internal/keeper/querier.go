package keeper

import (
	"bytes"
	"encoding/json"
	"github.com/ioioio8888/factio/x/factio/internal/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the factio Querier
const (
	QueryFact              = "getFact"
	QueryAddressDelegation = "getAddressDelegation"
	QueryFactList          = "getFactList"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryFact:
			return queryFact(ctx, path[1:], req, keeper)
		case QueryAddressDelegation:
			return queryAddressDelegation(ctx, path[1:], req, keeper)
		case QueryFactList:
			return queryFactList(ctx, req, keeper)
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

// return a list of title that the address has delegated to it
func queryAddressDelegation(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	var factDelegationList types.QueryResFactDelegationList
	var factDelegation types.FactDelegation
	iterator := keeper.GetFactDelegationIterator(ctx, path[0])

	for ; iterator.Valid(); iterator.Next() {
		if err := json.Unmarshal(iterator.Value(), &factDelegation); err != nil {
			panic(err)
		}
		factDelegationList = append(factDelegationList, factDelegation.Title)
	}
	res, err := codec.MarshalJSONIndent(keeper.cdc, factDelegationList)
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
