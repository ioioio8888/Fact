package factio

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/keeper"
)

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	//restore vote power every 10 blocks
	if ctx.BlockHeight()%10 == 0 {
		keeper.RestoreVotePower(ctx)
	}
}
