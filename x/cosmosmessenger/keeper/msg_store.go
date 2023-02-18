package keeper

import (
	"cosmos-messenger/x/cosmosmessenger/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreMessage(ctx sdk.Context, msg types.Message) {
	senderStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeySenderPrefix(msg.Sender)))
	receiverStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeyReceiverPrefix(msg.Receiver)))
	storedMsg := k.cdc.MustMarshal(&msg)
	senderStore.Set([]byte(msg.GetId()), storedMsg)
	receiverStore.Set([]byte(msg.GetId()), storedMsg)
	return
}

func GetMessageIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
