package global

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/config"
	"github.com/yann0917/check-in/utils"
)

var (
	Config     config.Server
	HttpClient *fiber.Agent
)

func init() {
	HttpClient = fiber.AcquireAgent()
}

func NewClient(cookie, referer string) *fiber.Agent {
	client := HttpClient.Debug()

	if cookies, err := utils.ParseCookiesMap(cookie); err == nil {
		for key, val := range cookies {
			client.Cookie(key, val)
		}
	}

	client.Referer(referer)
	client.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")

	return client
}
