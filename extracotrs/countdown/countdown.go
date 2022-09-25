package countdown

import (
	"fmt"
	"time"

	"github.com/yann0917/check-in/global"
	"github.com/yann0917/check-in/notification"
	"github.com/yann0917/check-in/utils"
)

func CalcDays() {
	content := ""
	list := global.Config.CountdownDays.List
	fmt.Printf("%#v\n", list)
	day, link := "", ""
	for k, v := range list {
		content += utils.Int2String(k+1) + ". 距离【<strong>" + v.Name + "</strong>】"
		targetTime := v.Date
		if time.Now().Before(targetTime) {
			link = "还有"
			day = utils.Float642String(targetTime.Sub(utils.TodayDate()).Hours()/24, 0)
		} else {
			// 正数日包含第一天 +1Day
			link = "已经第"
			day = utils.Float642String(utils.TodayDate().Add(24*time.Hour).Sub(targetTime).Hours()/24, 0)
		}
		content += link + "<strong>" + day + "</strong>" + "天"
		if len(v.Remark) > 0 {
			content += "【" + v.Remark + "】"
		}
		content += "\r\n"
	}
	if len(content) > 0 {
		notification.SendPushPlus("倒数日&纪念日提醒", content)
	}

}
