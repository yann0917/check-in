package notification

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/global"
)

type Param struct {
	Token       string `json:"token"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content"`
	Template    string `json:"template,omitempty"` // 消息模板 html,json,cloudMonitor 默认 html
	Topic       string `json:"topic,omitempty"`    // 群组编码，不填仅发送给自己；channel为webhook时无效
	Channel     string `json:"channel,omitempty"`  // 发送渠道, wechat,webhook,cp 默认 wechat
	Webhook     string `json:"webhook,omitempty"`
	CallbackUrl string `json:"callback_url,omitempty"` // 发送结果回调地址
}

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  string      `json:"data"`
	Count interface{} `json:"count"`
}

type CallbackResp struct {
	ShortCode  string `json:"shortCode"`  // 消息流水号
	SendStatus int    `json:"sendStatus"` // 发送状态；0未发送，1发送中，2发送成功，3发送失败
	Message    string `json:"message"`    // 推送错误内容（如有）
}

const url = "https://www.pushplus.plus/send"

func SendPushPlus(title, content string) {
	config := global.Config.Notification.PushPlus

	param := Param{
		Token:       config.Token,
		Title:       title,
		Content:     content,
		Template:    config.Template,
		CallbackUrl: config.Callback,
		Channel:     config.Channel,
		Webhook:     config.Webhook,
	}

	client := global.HttpClient.
		ConnectionClose().
		JSON(param)

	req := client.Request()
	req.Header.SetMethod(fiber.MethodPost)
	req.SetRequestURI(url)

	if err := client.Parse(); err != nil {
		panic(err)
	}

	var resp Response
	_, body, errs := client.Struct(&resp)
	fmt.Println(string(body))
	fmt.Println(errs)
}
