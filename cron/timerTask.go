package cron

import (
	"log"

	"github.com/yann0917/check-in/extracotrs/bilibili"
	"github.com/yann0917/check-in/extracotrs/countdown"
	"github.com/yann0917/check-in/extracotrs/juejin"
	"github.com/yann0917/check-in/extracotrs/tieba"
	"github.com/yann0917/check-in/pkg/timer"
)

var task timer.Timer

// documents see  https://pkg.go.dev/github.com/robfig/cron

// A cron expression represents a set of times, using 6 space-separated fields.
// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

// You may use one of several pre-defined schedules in place of a cron expression.
// Entry                  | Description                                | Equivalent To
// -----                  | -----------                                | -------------
// @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
// @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
// @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
// @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
// @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *

type timerTask struct {
	Name string
	Spec string
	Desc string
}

var taskList = []timerTask{
	{Name: "JueJinCheckIn", Spec: "0 1 8 * * *", Desc: "每天上午 08:01 执行掘金签到"},
	// {Name: "TieBaCheckIn", Spec: "0 0 10 * * *", Desc: "每天上午 10:00 执行贴吧签到"},
	{Name: "BiliCheckIn", Spec: "0 2 8 * * *", Desc: "每天上午 08:02 执行哔哩哔哩签到"},
	{Name: "CountdownDays", Spec: "0 0 6 * * *", Desc: "每天上午 06:00 执行纪念日提醒"},
}

func init() {
	task = timer.NewTimerTask()
}

func Task() {
	for _, t := range taskList {
		switch t.Name {
		case "JueJinCheckIn":
			_, err := task.AddTaskByFunc(t.Name, t.Spec, func() {
				juejin.CheckIn()
			})
			if err != nil {
				log.Print(err.Error())
			}
		case "TieBaCheckIn":
			_, err := task.AddTaskByFunc(t.Name, t.Spec, func() {
				tieba.SignAdd()
			})
			if err != nil {
				log.Print(err.Error())
			}
		case "BiliCheckIn":
			_, err := task.AddTaskByFunc(t.Name, t.Spec, func() {
				bilibili.CheckIn()
			})
			if err != nil {
				log.Print(err.Error())
			}
		case "CountdownDays":
			_, err := task.AddTaskByFunc(t.Name, t.Spec, func() {
				countdown.CalcDays()
			})
			if err != nil {
				log.Print(err.Error())
			}
		}

	}
}
