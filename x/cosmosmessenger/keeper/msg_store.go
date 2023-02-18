package keeper

import (
	"cosmos-messenger/x/cosmosmessenger/types"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) storeMessage(ctx sdk.Context, msg types.Message) {
	storedMsg := k.cdc.MustMarshal(&msg)
	fmt.Println(msg)
	k.getSenderStore(ctx, msg.SenderAddress).Set([]byte(msg.GetId()), storedMsg)
	k.getReceiverStore(ctx, msg.ReceiverAddress).Set([]byte(msg.GetId()), storedMsg)
	return
}

func (k Keeper) getSenderStore(ctx sdk.Context, senderAddress string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeySenderPrefix(senderAddress)))
}

func (k Keeper) getReceiverStore(ctx sdk.Context, receiverAddress string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyReceiverPrefix(receiverAddress)))
}

func (k Keeper) getSenderMessages(ctx sdk.Context, pagination *query.PageRequest, walletAddress string) ([]types.Message, *query.PageResponse, error) {
	return k.getMessages(pagination, k.getSenderStore(ctx, walletAddress))
}

func (k Keeper) getReceiverMessages(ctx sdk.Context, pagination *query.PageRequest, walletAddress string) ([]types.Message, *query.PageResponse, error) {
	return k.getMessages(pagination, k.getReceiverStore(ctx, walletAddress))
}

func (k Keeper) getMessages(pagination *query.PageRequest, store prefix.Store) ([]types.Message, *query.PageResponse, error) {
	var messages []types.Message
	paginatedRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var message types.Message
		if err := k.cdc.Unmarshal(value, &message); err != nil {
			return err
		}
		messages = append(messages, message)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return messages, paginatedRes, nil
}
