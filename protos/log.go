package protos

type UpdateLogConfigReq struct {
	Level int64 `json:"level"`
}

type LogConfig struct {
	Level int64 `json:"level"`
}
