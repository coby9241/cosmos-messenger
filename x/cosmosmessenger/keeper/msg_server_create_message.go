package keeper

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/segmentio/ksuid"
)

func (k msgServer) CreateMessage(goCtx context.Context, msg *types.MsgCreateMessage) (*types.MsgCreateMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// we allow sending to own wallet address fyi, just like slack allows it for instance
	senderKey, senderFound := k.GetEncryptionKey(ctx, msg.GetCreator())
	if !senderFound {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("sender wallet %s is not registered, please register the public key", msg.GetCreator()))
	}

	receiverKey, receiverFound := k.GetEncryptionKey(ctx, msg.GetReceiverWalletAddress())
	if !receiverFound {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("receiver wallet %s is not registered, please register the public key", msg.GetCreator()))
	}

	chatMsg := types.Message{
		Body:            msg.Body,
		Id:              ksuid.New().String(),
		SenderAddress:   msg.Creator,
		ReceiverAddress: msg.ReceiverWalletAddress,
	}

	encryptedSenderMsg, err := EncryptMessage(chatMsg, senderKey.PubKey)
	if err != nil {
		ctx.Logger().Error(fmt.Errorf("failed to encrypt message for wallet address: %s with error: %w", msg.GetCreator(), err).Error())
		return nil, status.Error(codes.Internal, "failed to encrypt message")
	}
	encryptedReceiverMsg, err := EncryptMessage(chatMsg, receiverKey.PubKey)
	if err != nil {
		ctx.Logger().Error(fmt.Errorf("failed to encrypt message for wallet address: %s with error: %w", msg.GetCreator(), err).Error())
		return nil, status.Error(codes.Internal, "failed to encrypt message")
	}

	k.storeSenderMessage(ctx, encryptedSenderMsg)
	k.storeReceiverMessage(ctx, encryptedReceiverMsg)

	return &types.MsgCreateMessageResponse{
		Id: chatMsg.GetId(),
	}, nil
}
