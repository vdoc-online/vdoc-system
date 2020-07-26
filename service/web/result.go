package web

//Result json wapper
type Result struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

//OK result
func OK(data interface{}) *Result {
	return &Result{Code: 0, Msg: "success", Data: data}
}
