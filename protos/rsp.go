package protos

type RspInfo struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data ...interface{}) *RspInfo {
	if len(data) == 0 {
		return &RspInfo{}
	}
	return &RspInfo{Data: data[0]}
}

func Error(err error) *RspInfo {
	return &RspInfo{Code: UnknownError.Code, Msg: err.Error()}
}
