package factio

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ioioio8888/factio/x/factio/internal/keeper"
)

func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {

	factDelegationList := keeper.GetAllFactDelegation(ctx)
	fcoin, _ := sdk.ParseCoins("1factcoin")
	if ctx.BlockHeight() > 1 {
		for _, delegation := range factDelegationList.Delegations {
			// delegation is the element from all delegations for where we are
			keeper.CoinKeeper.AddCoins(ctx, delegation.Delegator, fcoin)
		}
	}
}
