package common

const (
	Success      = "0"
	Error        = "-1"
	ParamError   = "-2"
	AccountExist = "-7"
	MailNotExist = "-8"

	/*
		token校验
	*/
	TokenError      = "-3" // 生成token错误
	AuthBlank       = "-4" // 请求头auth为空
	AuthFormatError = "-5" // 请求头auth格式错误
	InvalidToken    = "-6" // 无效的token

	/*
		服务器错误
	*/
	ServerError = "-1000"
)

type BaseResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BadResponse(code string) (baseResponse *BaseResponse) {
	baseResponse = new(BaseResponse)
	baseResponse.Code = code
	return
}
