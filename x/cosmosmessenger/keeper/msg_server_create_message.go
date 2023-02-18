package keeper

import (
	"context"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) CreateMessage(goCtx context.Context, msg *types.MsgCreateMessage) (*types.MsgCreateMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator == msg.ReceiverWalletAddress {
		return nil, status.Error(codes.InvalidArgument, "cannot send messages to own wallet address")
	}

	chatMsg := types.Message{
		Body:            msg.Body,
		Id:              ksuid.New().String(),
		SenderAddress:   msg.Creator,
		ReceiverAddress: msg.ReceiverWalletAddress,
	}
	k.storeMessage(ctx, chatMsg)
	ctx.Logger().Info(msg.ReceiverWalletAddress)
	ctx.Logger().Info(msg.Creator)

	return &types.MsgCreateMessageResponse{
		Id: chatMsg.GetId(),
	}, nil
}
