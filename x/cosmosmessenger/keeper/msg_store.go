package keeper

import (
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) storeMessage(ctx sdk.Context, msg types.Message) {
	storedMsg := k.cdc.MustMarshal(&msg)
	k.getSenderStore(ctx, msg.Sender).Set([]byte(msg.GetId()), storedMsg)
	k.getReceiverStore(ctx, msg.Receiver).Set([]byte(msg.GetId()), storedMsg)
	return
}

func (k Keeper) getSenderStore(ctx sdk.Context, sender string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeySenderPrefix(sender)))
}

func (k Keeper) getReceiverStore(ctx sdk.Context, receiver string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyReceiverPrefix(receiver)))
}

func (k Keeper) getSenderMessages(ctx sdk.Context, pagination *query.PageRequest, user string) ([]types.Message, *query.PageResponse, error) {
	var msgs []types.Message

	paginatedRes, err := query.Paginate(k.getSenderStore(ctx, user), pagination, func(key []byte, value []byte) error {
		var msg types.Message
		if err := k.cdc.Unmarshal(value, &msg); err != nil {
			return err
		}
		msgs = append(msgs, msg)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return msgs, paginatedRes, nil
}
