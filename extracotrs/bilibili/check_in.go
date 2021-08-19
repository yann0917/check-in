package bilibili

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

func NewClient() *fiber.Agent {
	return global.NewClient(global.Config.Bilibili.Cookie, referer)
}

func CheckIn() {

	client := NewClient()
	req := client.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(host + apis["check_in"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	resp.Data = new(DoSignData)

	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	// fmt.Println(resp.Data)
	if len(errs) != 0 {
		notification.SendPushPlus("【"+appName+"】签到失败", errs[0].Error())
		return
	}
	data, ok := resp.Data.(*DoSignData)

	if ok && resp.Code == 0 {
		content := data.Text
		notification.SendPushPlus("【"+appName+"】签到成功", content)
	} else {
		content := resp.Message
		switch resp.Code {
		case -101:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}
