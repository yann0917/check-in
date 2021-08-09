package notification

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/global"
)

type Param struct {
	Token    string `json:"token"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content"`
	Template string `json:"template,omitempty"` // 消息模板 html,json,cloudMonitor 默认 html
}

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  string      `json:"data"`
	Count interface{} `json:"count"`
}

const url = "http://pushplus.hxtrip.com/send"

type Client struct {
	Client *fiber.Agent
}

func SendPushPlus(title, content string) {
	param := Param{
		Token:    global.Config.Notification.PushPlus.Token,
		Title:    title,
		Content:  content,
		Template: "html",
	}

	client := global.HttpClient.JSON(param)

	req := client.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(url)

	fmt.Println(string(req.Body()))
	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
	fmt.Println(resp.Data)
}
