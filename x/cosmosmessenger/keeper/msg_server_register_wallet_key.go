package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterWalletKey(goCtx context.Context, msg *types.MsgRegisterWalletKey) (*types.MsgRegisterWalletKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, err := stringToPublicKey(msg.GetPubkey()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "please provide a valid base64 encoded RSA public keys")
	}

	k.StoreEncryptionKey(ctx, types.EncryptKey{
		PubKey: msg.GetPubkey(),
	}, msg.GetCreator())

	return &types.MsgRegisterWalletKeyResponse{Success: true}, nil
}
