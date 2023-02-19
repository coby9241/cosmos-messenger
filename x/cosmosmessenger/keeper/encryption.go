package keeper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmos-messenger/x/cosmosmessenger/types"
)

func EncryptMessage(msg types.Message, pubKey string) (types.Message, error) {
	pub, err := stringToPublicKey(pubKey)
	if err != nil {
		return msg, err
	}

	encryptedMsg, err := encryptWithPublicKey([]byte(msg.Body), pub)
	if err != nil {
		return msg, err
	}

	encodedMsg := base64.StdEncoding.EncodeToString(encryptedMsg)

	msg.Body = encodedMsg
	return msg, nil
}

// EncryptWithPublicKey encrypts data with public key
func encryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, msg)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// StringToPublicKey converts string to public key
func stringToPublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	decodePubKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid base64 encoded public key provided")
	}

	block, _ := pem.Decode(decodePubKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub.(type) {
	case *rsa.PublicKey:
		return pub.(*rsa.PublicKey), nil
	default:
		return nil, errors.New("unsupported key type")
	}
}
