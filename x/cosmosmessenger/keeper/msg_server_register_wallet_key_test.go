package keeper_test

import (
	"testing"

	"cosmos-messenger/testutil/sample"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/stretchr/testify/require"
)

func TestMsgServer_RegisterWalletKey(t *testing.T) {
	creatorAddr := sample.AccAddress()
	_, pubKeyStr := getEncryptionKeys(t)

	t.Run("should register new pub key with no error", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		// act
		resp, err := msgSvr.RegisterWalletKey(ctx, &types.MsgRegisterWalletKey{
			Creator: creatorAddr,
			Pubkey:  pubKeyStr,
		})
		// assert
		require.NoError(t, err)
		require.True(t, resp.GetSuccess())
	})

	t.Run("should fail to register new pub key if key is not valid", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		// act
		_, err := msgSvr.RegisterWalletKey(ctx, &types.MsgRegisterWalletKey{
			Creator: creatorAddr,
			Pubkey:  "example",
		})
		// assert
		require.Error(t, err)
	})
}
