package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"testing"

	"cosmos-messenger/testutil/sample"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_ShowReceivedMessages(t *testing.T) {
	senderAddr := sample.AccAddress()
	receiverAddr := sample.AccAddress()

	t.Run("should show received message after sending a message", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, k, ctx := setupMsgServer(t)
		// act
		createRes, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)

		res, err := k.ShowReceivedMessages(ctx, &types.QueryShowReceivedMessagesRequest{
			WalletAddress: receiverAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(res.Messages))
		require.Equal(t, senderAddr, res.Messages[0].SenderAddress)
		require.Equal(t, receiverAddr, res.Messages[0].ReceiverAddress)
		require.Equal(t, "this is a test message", res.Messages[0].Body)
		require.Equal(t, createRes.GetId(), res.Messages[0].GetId())
	})

	t.Run("should not show message when calling ShowReceivedMessages if requested wallet is sender", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, k, ctx := setupMsgServer(t)
		// act
		_, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               receiverAddr,
			ReceiverWalletAddress: senderAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)

		res, err := k.ShowReceivedMessages(ctx, &types.QueryShowReceivedMessagesRequest{
			WalletAddress: receiverAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 0, len(res.Messages))
	})

	t.Run("should show all received messages", func(t *testing.T) {
		// arrange
		var err error
		msgSvr, k, ctx := setupMsgServer(t)
		messages := []string{
			"message 1",
			"message 2",
			"message 3",
		}
		// act
		for _, msg := range messages {
			_, err = msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
				Creator:               senderAddr,
				ReceiverWalletAddress: receiverAddr,
				Body:                  msg,
			})
			require.NoError(t, err)
		}

		res, err := k.ShowReceivedMessages(ctx, &types.QueryShowReceivedMessagesRequest{
			WalletAddress: receiverAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 3, len(res.Messages))
	})

	t.Run("should show all received messages with pagination", func(t *testing.T) {
		t.Parallel()
		// arrange
		var err error
		msgSvr, k, ctx := setupMsgServer(t)
		messages := []string{
			"message 1",
			"message 2",
			"message 3",
		}
		// act
		for _, msg := range messages {
			_, err = msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
				Creator:               senderAddr,
				ReceiverWalletAddress: receiverAddr,
				Body:                  msg,
			})
			require.NoError(t, err)
		}

		res, err := k.ShowReceivedMessages(ctx, &types.QueryShowReceivedMessagesRequest{
			WalletAddress: receiverAddr,
			Pagination:    &query.PageRequest{Limit: 1},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res.Pagination)
		require.Equal(t, 1, len(res.Messages))

		nextPageRes, err := k.ShowReceivedMessages(ctx, &types.QueryShowReceivedMessagesRequest{
			WalletAddress: receiverAddr,
			Pagination:    &query.PageRequest{Key: res.Pagination.GetNextKey()},
		})
		require.NoError(t, err)
		require.Empty(t, nextPageRes.Pagination)
		require.Equal(t, 2, len(nextPageRes.Messages))
	})
}

func TestKeeper_ShowSentMessages(t *testing.T) {
	senderAddr := sample.AccAddress()
	receiverAddr := sample.AccAddress()

	t.Run("should show sent message after sending a message", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, k, ctx := setupMsgServer(t)
		// act
		createRes, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               senderAddr,
			ReceiverWalletAddress: receiverAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)

		res, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(res.Messages))
		require.Equal(t, senderAddr, res.Messages[0].SenderAddress)
		require.Equal(t, receiverAddr, res.Messages[0].ReceiverAddress)
		require.Equal(t, "this is a test message", res.Messages[0].Body)
		require.Equal(t, createRes.GetId(), res.Messages[0].GetId())
	})

	t.Run("should not show message when calling ShowReceivedMessages if requested wallet is sender", func(t *testing.T) {
		t.Parallel()
		// arrange
		msgSvr, k, ctx := setupMsgServer(t)
		// act
		_, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               receiverAddr,
			ReceiverWalletAddress: senderAddr,
			Body:                  "this is a test message",
		})
		// assert
		require.NoError(t, err)

		res, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 0, len(res.Messages))
	})

	t.Run("should show all sent messages", func(t *testing.T) {
		// arrange
		var err error
		msgSvr, k, ctx := setupMsgServer(t)
		messages := []string{
			"message 1",
			"message 2",
			"message 3",
		}
		// act
		for _, msg := range messages {
			_, err = msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
				Creator:               senderAddr,
				ReceiverWalletAddress: receiverAddr,
				Body:                  msg,
			})
			require.NoError(t, err)
		}

		res, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 3, len(res.Messages))
	})

	t.Run("should show all sent messages with pagination", func(t *testing.T) {
		t.Parallel()
		// arrange
		var err error
		msgSvr, k, ctx := setupMsgServer(t)
		messages := []string{
			"message 1",
			"message 2",
			"message 3",
		}
		// act
		for _, msg := range messages {
			_, err = msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
				Creator:               senderAddr,
				ReceiverWalletAddress: receiverAddr,
				Body:                  msg,
			})
			require.NoError(t, err)
		}

		res, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
			Pagination:    &query.PageRequest{Limit: 1},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res.Pagination)
		require.Equal(t, 1, len(res.Messages))

		nextPageRes, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: senderAddr,
			Pagination:    &query.PageRequest{Key: res.Pagination.GetNextKey()},
		})
		require.NoError(t, err)
		require.Empty(t, nextPageRes.Pagination)
		require.Equal(t, 2, len(nextPageRes.Messages))
	})
}

func TestKeeper_ShowMixedMessages(t *testing.T) {
	walletAddr := sample.AccAddress()

	t.Run("should show a message in both ShowSentMessages and ShowReceivedMessages if sender == receiver", func(t *testing.T) {
		msgSvr, k, ctx := setupMsgServer(t)
		res, err := msgSvr.CreateMessage(ctx, &types.MsgCreateMessage{
			Creator:               walletAddr,
			ReceiverWalletAddress: walletAddr,
			Body:                  "test sending to myself",
		})
		require.NoError(t, err)

		// check sent msgs for wallet
		querySender, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: walletAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(querySender.Messages))
		require.Equal(t, walletAddr, querySender.Messages[0].SenderAddress)
		require.Equal(t, walletAddr, querySender.Messages[0].ReceiverAddress)
		require.Equal(t, "test sending to myself", querySender.Messages[0].Body)
		require.Equal(t, res.GetId(), querySender.Messages[0].GetId())

		// check received msgs for wallet
		queryReceiver, err := k.ShowSentMessages(ctx, &types.QueryShowSentMessagesRequest{
			WalletAddress: walletAddr,
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(queryReceiver.Messages))
		require.Equal(t, walletAddr, queryReceiver.Messages[0].SenderAddress)
		require.Equal(t, walletAddr, queryReceiver.Messages[0].ReceiverAddress)
		require.Equal(t, "test sending to myself", queryReceiver.Messages[0].Body)
		require.Equal(t, res.GetId(), queryReceiver.Messages[0].GetId())
	})
}
