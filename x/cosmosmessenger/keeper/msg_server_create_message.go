package keeper

import (
	"context"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/segmentio/ksuid"
)

func (k msgServer) CreateMessage(goCtx context.Context, msg *types.MsgCreateMessage) (*types.MsgCreateMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// we allow sending to own wallet address fyi, just like slack allows it for instance
	chatMsg := types.Message{
		Body:            msg.Body,
		Id:              ksuid.New().String(),
		SenderAddress:   msg.Creator,
		ReceiverAddress: msg.ReceiverWalletAddress,
	}
	k.storeMessage(ctx, chatMsg)

	return &types.MsgCreateMessageResponse{
		Id: chatMsg.GetId(),
	}, nil
}
