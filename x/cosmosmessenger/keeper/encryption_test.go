package keeper_test

import (
	"crypto/rsa"
	"testing"

	"cosmos-messenger/testutil/encryption"
	. "cosmos-messenger/x/cosmosmessenger/keeper"
	"cosmos-messenger/x/cosmosmessenger/types"
	"github.com/stretchr/testify/require"
)

func getEncryptionKeys(t *testing.T) (*rsa.PrivateKey, string) {
	// load public key from testdata
	privateKey, pubKey, err := encryption.SetupTestKeys(2048)
	require.NoError(t, err)
	pubKeyStr, err := encryption.ExportRsaPublicKeyAsPemStr(pubKey)
	require.NoError(t, err)
	return privateKey, pubKeyStr
}

func TestEncryptSenderMessage(t *testing.T) {
	privateKey, pubKeyStr := getEncryptionKeys(t)
	t.Run("should encrypt message and base64 encode successfully, and can be decrypted by private key", func(t *testing.T) {
		t.Parallel()
		res, err := EncryptMessage(types.Message{
			Body: "example",
		}, pubKeyStr)
		require.NoError(t, err)
		decryptedMsg, err := encryption.DecryptMessageGivenKey(res.Body, privateKey)
		require.NoError(t, err)
		require.Equal(t, "example", string(decryptedMsg))
	})

	t.Run("should fail to encrypt if public key is invalid", func(t *testing.T) {
		t.Parallel()
		_, err := EncryptMessage(types.Message{
			Body: "example",
		}, "wrong key")
		require.Error(t, err)
	})

	t.Run("should fail to decrypt if public key/private key does not match", func(t *testing.T) {
		t.Parallel()
		// load public key from testdata
		newPrivateKey, _, err := encryption.SetupTestKeys(2048)
		require.NoError(t, err)

		res, err := EncryptMessage(types.Message{
			Body: "example",
		}, pubKeyStr)
		require.NoError(t, err)
		_, err = encryption.DecryptMessageGivenKey(res.Body, newPrivateKey)
		require.Error(t, err)
	})
}
