package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateFact{}, "factio/CreateFact", nil)
	cdc.RegisterConcrete(MsgEditFact{}, "factio/EditFact", nil)
	cdc.RegisterConcrete(MsgDelegateFact{}, "factio/DelegateFact", nil)
	cdc.RegisterConcrete(MsgUnDelegateFact{}, "factio/UndelegateFact", nil)
	cdc.RegisterConcrete(MsgVoteFact{}, "factio/VoteFact", nil)

}
