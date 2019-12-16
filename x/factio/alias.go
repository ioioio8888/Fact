package factio

import (
	"github.com/ioioio8888/factio/x/factio/internal/keeper"
	"github.com/ioioio8888/factio/x/factio/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
	Fact   = types.Fact
)
