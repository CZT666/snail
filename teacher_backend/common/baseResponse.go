package common

const (
	Success    = "0"
	Error      = "-1"
	ParamError = "-2"
)

type BaseResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
