package comm

import "errors"

var (
	internalServerError = errors.New("internal server error")
	base64DecodeFailed  = errors.New("base64 decode failed")
	pemDecodeFailed     = errors.New("pem decode failed")
	aesInvalidKey       = errors.New("aes invalid key")
	aesInvalidIV        = errors.New("aes invalid iv")
	aesDecryptFailed    = errors.New("aes decrypt failed")
	rsaInvalidPubKey    = errors.New("rsa invalid public key")
	rsaInvalidPvtKey    = errors.New("rsa invalid private key")
	rsaDecryptFailed    = errors.New("rsa decrypt failed")
)
