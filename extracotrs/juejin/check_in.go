package juejin

import (
	"log"
	"strconv"

	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
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

func NewClient() *global.Client {
	return global.NewClient(global.Config.JueJin.Cookie, referer)
}

func GetCheckInStatus() bool {
	client := NewClient()
	resp, err := client.Get(host+apis["today_status"], client.Headers)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
		return false
	}

	var result Response
	err = resp.ToJSON(&result)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(result)

	status, ok := result.Data.(bool)
	if ok && status {
		notification.SendPushPlus("【"+appName+"】签到状态", "您今天已经签过到了")
	}

	return status
}

func CheckIn() {

	// status := GetCheckInStatus()
	// if status {
	// 	return
	// }

	client := NewClient()
	resp, err := client.Post(host+apis["check_in"], client.Headers)
	if err != nil {
		notification.SendPushPlus("【"+appName+"】签到失败", err.Error())
		return
	}
	var result Response
	result.Data = new(CheckInData)
	err = resp.ToJSON(&result)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)

	data, ok := result.Data.(*CheckInData)

	if ok && result.ErrNo == 0 {
		content := "新增矿石: " + strconv.Itoa(data.IncrPoint) + "，当前矿石数: " + strconv.Itoa(data.SumPoint)
		notification.SendPushPlus("【"+appName+"】签到成功", content)
		LotteryDraw()
	} else {
		content := result.ErrMsg
		switch result.ErrNo {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】签到失败", content)
	}
}

func LotteryDraw() {
	client := NewClient()
	resp, err := client.Post(host+apis["lottery_draw"], client.Headers)
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

	log.Println(result)

	data, ok := result.Data.(*LotteryDrawData)

	if ok && result.ErrNo == 0 {
		content := "奖品: " + data.LotteryName
		notification.SendPushPlus("【"+appName+"】抽奖成功", content)
	} else {
		content := result.ErrMsg
		switch result.ErrNo {
		case 403:
			content += "，请更新 cookie 后再执行。"
		}
		notification.SendPushPlus("【"+appName+"】抽奖失败", content)
	}
}
