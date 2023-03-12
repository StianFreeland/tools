package protos

type GetPwdReq struct {
	Length int `json:"length" binding:"required"`
}

type PwdInfo struct {
	Pwd string `json:"pwd"`
}

type GetKeyReq struct {
	Length int `json:"length" binding:"required"`
}

type KeyInfo struct {
	PvtKey string `json:"pvtKey"`
	PubKey string `json:"pubKey"`
}

type MD5EncryptReq struct {
	Msg string `json:"msg" binding:"required"`
}

type MD5EncryptedInfo struct {
	Msg string `json:"msg"`
}

type SHAEncryptReq struct {
	Msg string `json:"msg" binding:"required"`
}

type SHAEncryptedInfo struct {
	Msg string `json:"msg"`
}

type AESEncryptReq struct {
	Msg string `json:"msg" binding:"required"`
	Key string `json:"key" binding:"required"`
	IV  string `json:"iv" binding:"required"`
}

type AESDecryptReq struct {
	Msg string `json:"msg" binding:"required"`
	Key string `json:"key" binding:"required"`
	IV  string `json:"iv" binding:"required"`
}

type AESEncryptedInfo struct {
	Msg string `json:"msg"`
}

type AESDecryptedInfo struct {
	Msg string `json:"msg"`
}

type RSAEncryptReq struct {
	Msg string `json:"msg" binding:"required"`
	Key string `json:"key" binding:"required"`
}

type RSADecryptReq struct {
	Msg string `json:"msg" binding:"required"`
	Key string `json:"key" binding:"required"`
}

type RSAEncryptedInfo struct {
	Msg string `json:"msg"`
}

type RSADecryptedInfo struct {
	Msg string `json:"msg"`
}
