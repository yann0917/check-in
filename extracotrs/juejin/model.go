package juejin

type Response struct {
	ErrNo  int         `json:"err_no"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

type CheckInData struct {
	IncrPoint int `json:"incr_point"`
	SumPoint  int `json:"sum_point"`
}
