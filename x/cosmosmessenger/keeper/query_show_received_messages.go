package keeper

import (
	"context"

	"cosmos-messenger/x/cosmosmessenger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowReceivedMessages(goCtx context.Context, req *types.QueryShowReceivedMessagesRequest) (*types.QueryShowReceivedMessagesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	messages, paginatedRes, err := k.getReceiverMessages(ctx, req.GetPagination(), req.User)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowReceivedMessagesResponse{
		Message:    messages,
		Pagination: paginatedRes,
	}, nil
}
