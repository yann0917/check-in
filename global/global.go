package global

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yann0917/check-in/config"
)

var (
	Config     config.Server
	HttpClient *fiber.Agent
)

func init() {
	HttpClient = fiber.AcquireAgent()
}
