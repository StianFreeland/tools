package protos

var (
	InternalServerError = &RspInfo{Code: 10000, Msg: "Internal Server Error"}
)

var (
	InvalidRequestParameters = &RspInfo{Code: 20000, Msg: "Invalid Request Parameters"}
)

var (
	InvalidKeyLength = &RspInfo{Code: 30000, Msg: "Invalid Key Length"}
)

var (
	UnknownError = &RspInfo{Code: 90000, Msg: "Unknown Error"}
)
