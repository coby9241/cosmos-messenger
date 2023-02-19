package keeper_test

import (
	"context"
	"crypto/rsa"
	"regexp"
	"testing"

	"cosmos-messenger/testutil/encryption"
	"cosmos-messenger/testutil/sample"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/stretchr/testify/require"
)

func IsValidKsuid(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]{27}$")
	return re.MatchString(str)
}

func registerWalletForUser(ctx context.Context, t *testing.T, msgSvr types.MsgServer, walletAddress string) (*rsa.PrivateKey, string) {
	privateKey, pubKeyStr := getEncryptionKeys(t)

	_, err := msgSvr.RegisterWalletKey(ctx, &types.MsgRegisterWalletKey{
		Creator: walletAddress,
		Pubkey:  pubKeyStr,
	})
	require.NoError(t, err)
	return privateKey, pubKeyStr
}

func TestMsgServer_CreateMessage(t *testing.T) {
	senderAddr := sample.AccAddress()
	receiverAddr := sample.AccAddress()

	t.Run("should create new message with no error and with valid KSUID", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		_, _ = registerWalletForUser(ctx, t, msgSvr, senderAddr)
		_, _ = registerWalletForUser(ctx, t, msgSvr, receiverAddr)
		// act
		resp, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)
		require.True(t, IsValidKsuid(resp.GetId()))
	})

	t.Run("should create new message and verify message body from store", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, k, ctx := setupMsgServer(t)
		privateKey, _ := registerWalletForUser(ctx, t, msgSvr, senderAddr)
		_, _ = registerWalletForUser(ctx, t, msgSvr, receiverAddr)
		// act
		resp, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)
		require.True(t, IsValidKsuid(resp.GetId()))

		result, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
		})
		// assert
		require.NoError(t, err)
		require.Equal(t, 1, len(result.Messages))
		require.Equal(t, senderAddr, result.Messages[0].SenderAddress)
		require.Equal(t, receiverAddr, result.Messages[0].ReceiverAddress)

		decryptedMsg, err := encryption.DecryptMessageGivenKey(result.Messages[0].Body, privateKey)
		require.NoError(t, err)
		require.Equal(t, "this is a test message", string(decryptedMsg))
	})

	t.Run("should create multiple messages with no issues", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		_, _ = registerWalletForUser(ctx, t, msgSvr, senderAddr)
		_, _ = registerWalletForUser(ctx, t, msgSvr, receiverAddr)
		messages := []string{
			"test message 1",
			"test message 2",
			"test message 3",
		}
		// act
		for _, msg := range messages {
			resp, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
				Creator:               senderAddr,
				ReceiverWalletAddress: receiverAddr,
				Body:                  msg,
			})
			// assert
			require.NoError(t, err)
			require.True(t, IsValidKsuid(resp.GetId()))
		}
	})

	t.Run("should fail to create message if sender is not registered", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		// act
		_, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.Error(t, err)
	})

	t.Run("should fail to create message if receiver is not registered", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, _, ctx := setupMsgServer(t)
		_, _ = registerWalletForUser(ctx, t, msgSvr, senderAddr)
		// act
		_, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.Error(t, err)
	})
}
