package loggers

import "github.com/gofiber/fiber/v2"

type ILoggerMiddleware interface {
	Log(ctx *fiber.Ctx) error
}
