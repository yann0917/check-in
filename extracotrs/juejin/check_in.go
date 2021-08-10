package juejin

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
	"github.com/yann0917/check-in/utils"
)

const (
	appName = "掘金"
	host    = "https://api.juejin.cn"
	referer = "https://juejin.cn/"
)

var apis = map[string]string{
	"today_status": "/growth_api/v1/get_today_status",
	"lottery_draw": "/growth_api/v1/lottery/draw",
	"check_in":     "/growth_api/v1/check_in" + "?_signature=" + global.Config.JueJin.Signature,
}

type Client struct {
	Client *fiber.Agent
}

func NewClient() *fiber.Agent {
	client := global.HttpClient

	if cookies, err := utils.ParseCookiesMap(global.Config.JueJin.Cookie); err == nil {
		for key, val := range cookies {
			client.Cookie(key, val)
		}
	}

	client.ContentType("application/json")
	client.Referer(referer)
	client.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

	return client
}
func GetCheckInStatus() bool {
	client := NewClient()
	req := client.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(host + apis["today_status"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	status, ok := resp.Data.(bool)
	if ok && status {
		notification.SendPushPlus("【"+appName+"】签到状态", "您今天已经签过到了")
	}

	return status
}

func CheckIn() {

	status := GetCheckInStatus()
	if status {
		return
	}

	client := NewClient()
	req := client.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(host + apis["check_in"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	resp.Data = new(CheckInData)

	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	// fmt.Println(resp.Data)
	data, ok := resp.Data.(*CheckInData)

	if ok && resp.ErrNo == 0 {
		content := "新增矿石: " + strconv.Itoa(data.IncrPoint) + "，当前矿石数: " + strconv.Itoa(data.SumPoint)
		notification.SendPushPlus("【"+appName+"】签到成功", content)
		LotteryDraw()
	} else {
		content := resp.ErrMsg
		switch resp.ErrNo {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}

func LotteryDraw() {
	client := NewClient()
	req := client.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(host + apis["lottery_draw"])

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	resp.Data = new(LotteryDrawData)

	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	// fmt.Println(resp.Data)
	data, ok := resp.Data.(*LotteryDrawData)

	if ok && resp.ErrNo == 0 {
		content := "奖品: " + data.LotteryName
		notification.SendPushPlus("【"+appName+"】抽奖成功", content)
	} else {
		content := resp.ErrMsg
		switch resp.ErrNo {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】抽奖失败", content)
	}
}
