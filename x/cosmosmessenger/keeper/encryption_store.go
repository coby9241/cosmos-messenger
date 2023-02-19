package keeper

import (
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) StoreEncryptionKey(ctx sdk.Context, key types.EncryptKey, walletAddr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptionKey))
	storedKeys := k.cdc.MustMarshal(&key)
	store.Set([]byte(walletAddr), storedKeys)
	return
}

func (k Keeper) GetEncryptionKey(ctx sdk.Context, walletAddr string) (types.EncryptKey, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptionKey))
	res := store.Get([]byte(walletAddr))
	if res == nil {
		return types.EncryptKey{}, false
	}

	var encryptKey types.EncryptKey
	k.cdc.MustUnmarshal(res, &encryptKey)
	return encryptKey, true
}
