package bilibili

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Ttl     int         `json:"ttl"`
	Data    interface{} `json:"data"`
}

type DoSignData struct {
	Text        string `json:"text"`
	SpecialText string `json:"specialText"`
	AllDays     int    `json:"allDays"`
	HadSignDays int    `json:"hadSignDays"`
	IsBonusDay  int    `json:"isBonusDay"`
}
