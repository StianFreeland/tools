package comm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AESEncryptMsg(msg string, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", aesInvalidKey
	}
	if len(iv) != block.BlockSize() {
		return "", aesInvalidIV
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	padding := block.BlockSize() - len(msg)%block.BlockSize()
	padded := append([]byte(msg), bytes.Repeat([]byte{byte(padding)}, padding)...)
	encrypted := make([]byte, len(padded))
	blockMode.CryptBlocks(encrypted, padded)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func AESDecryptMsg(msg string, key []byte, iv []byte) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", base64DecodeFailed
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", aesInvalidKey
	}
	if len(iv) != block.BlockSize() {
		return "", aesInvalidIV
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	padded := make([]byte, len(decoded))
	blockMode.CryptBlocks(padded, decoded)
	length := len(padded)
	if length == 0 {
		return "", aesDecryptFailed
	}
	padding := int(padded[length-1])
	if length < padding {
		return "", aesDecryptFailed
	}
	return string(padded[:length-padding]), nil
}
