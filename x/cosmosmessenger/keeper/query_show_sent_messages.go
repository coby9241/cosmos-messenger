package keeper

import (
	"context"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowSentMessages(goCtx context.Context, req *types.QueryShowSentMessagesRequest) (*types.QueryShowSentMessagesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	msgs, paginatedRes, err := k.getSenderMessages(ctx, req.GetPagination(), req.User)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowSentMessagesResponse{
		Message:    msgs,
		Pagination: paginatedRes,
	}, nil
}
