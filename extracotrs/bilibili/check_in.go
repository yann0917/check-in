package bilibili

import (
	"log"
	"strconv"

	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
)

const (
	appName = "哔哩哔哩"
	host    = "https://api.live.bilibili.com"
	referer = "https://live.bilibili.com/"
)

var apis = map[string]string{
	"check_in": "/sign/doSign",
}

func NewClient() *global.Client {
	return global.NewClient(global.Config.Bilibili.Cookie, referer)
}

func CheckIn() {

	client := NewClient()
	resp, err := client.Get(host+apis["check_in"], client.Headers)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
		return
	}
	var result Response
	err = resp.ToJSON(&result)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.Dump())

	log.Println(result)
	//
	// req := client.Request()
	// req.Header.SetMethod(fiber.MethodGet)
	// req.SetRequestURI(host + apis["check_in"])
	//
	// if err := client.Parse(); err != nil {
	// 	panic(err)
	// }
	//
	// var resp Response
	// resp.Data = new(DoSignData)
	//
	// _, body, errs := client.Struct(&resp)
	// fmt.Println(string(body))
	// fmt.Println(errs)
	// fmt.Println(resp.Data)
	//

	data, ok := result.Data.(*DoSignData)

	if ok && result.Code == 0 {
		content := "已签到" + strconv.Itoa(data.HadSignDays) + "天，"
		if data.IsBonusDay == 1 {
			content += "礼物签到，" + data.Text
		} else {
			content += data.Text
		}

		if data.SpecialText != "" {
			content += data.SpecialText
		}

		notification.SendPushPlus("【"+appName+"】签到成功", content)
	} else {
		content := result.Message
		switch result.Code {
		case -101:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}
