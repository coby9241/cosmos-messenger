package keeper

import (
	"context"
	"github.com/segmentio/ksuid"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMessage(goCtx context.Context, msg *types.MsgCreateMessage) (*types.MsgCreateMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	chatMsg := types.Message{
		Sender:   msg.Creator,
		Receiver: msg.ReceiverWalletAddress,
		Body:     msg.Body,
		Id:       ksuid.New().String(),
	}
	k.storeMessage(ctx, chatMsg)
	ctx.Logger().Info(msg.ReceiverWalletAddress)
	ctx.Logger().Info(msg.Creator)

	return &types.MsgCreateMessageResponse{
		Id: chatMsg.GetId(),
	}, nil
}
