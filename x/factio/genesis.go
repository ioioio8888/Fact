package factio

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/factio/x/factio/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	FactRecords []Fact `json:"facts_records"`
}

func NewGenesisState(whoIsRecords []Fact) GenesisState {
	return GenesisState{FactRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.FactRecords {
		if record.Creator == nil {
			return fmt.Errorf("invalid factsRecord: Value: . Error: Missing Owner")
		}
		if record.Description == "" {
			return fmt.Errorf("invalid factsRecord: Owner: . Error: Missing Value")
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: . Error: Missing Price")
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		FactRecords: []Fact{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.FactRecords {
		keeper.SetFact(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Fact
	iterator := k.GetFactIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		// titlewithprefix := string(iterator.Key())
		title := bytes.Trim(iterator.Key(), string(types.FactKey))
		fact := k.GetFact(ctx, string(title))
		records = append(records, fact)

	}
	return GenesisState{FactRecords: records}
}
