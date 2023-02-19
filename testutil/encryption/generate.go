package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func SetupTestKeys(testKeySize int) (priKey *rsa.PrivateKey, pubKey *rsa.PublicKey, err error) {
	if priKey, err = rsa.GenerateKey(rand.Reader, testKeySize); err != nil {
		return
	}

	pubKey = &priKey.PublicKey
	return
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return base64.URLEncoding.EncodeToString(pubkeyPem), nil
}
