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

type LotteryDrawData struct {
	Id           int    `json:"id"`
	LotteryId    string `json:"lottery_id"`
	LotteryName  string `json:"lottery_name"`
	LotteryType  int    `json:"lottery_type"`
	LotteryImage string `json:"lottery_image"`
	HistoryId    string `json:"history_id"`
}
