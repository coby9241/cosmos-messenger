package keeper_test

import (
	"context"
	"testing"

	keepertest "cosmos-messenger/testutil/keeper"
	"cosmos-messenger/x/cosmosmessenger/keeper"
	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CosmosmessengerKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
