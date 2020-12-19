package vo

type ResetPwdRequest struct {
	Mail  string `json:"mail"`
	Proof string `json:"proof"`
}
