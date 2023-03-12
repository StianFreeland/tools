package comm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func RSAEncryptMsg(msg string, key []byte) (string, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return "", pemDecodeFailed
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", rsaInvalidPubKey
	}
	pk, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return "", rsaInvalidPubKey
	}
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pk, []byte(msg), []byte{})
	if err != nil {
		return "", internalServerError
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func RSADecryptMsg(msg string, key []byte) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", base64DecodeFailed
	}
	block, _ := pem.Decode(key)
	if block == nil {
		return "", pemDecodeFailed
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", rsaInvalidPvtKey
	}
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, pk, decoded, []byte{})
	if err != nil {
		return "", rsaDecryptFailed
	}
	return string(decrypted), nil
}
