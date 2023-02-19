package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func DecryptMessageGivenKey(encodedString string, privateKey *rsa.PrivateKey) ([]byte, error) {
	decodedStr, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	decMsg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedStr)
	if err != nil {
		return nil, err
	}

	return decMsg, nil
}

func LoadKeyFromPemFile(pemFilepath string) (*rsa.PrivateKey, error) {
	b, err := os.ReadFile(pemFilepath) // just pass the file name
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
