package keeper

import (
	"context"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMessage(goCtx context.Context, msg *types.MsgCreateMessage) (*types.MsgCreateMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMessageResponse{}, nil
}
