package request_timeouts

import "github.com/gofiber/fiber/v2"

type IRequestTimeoutDMiddleware interface {
	RequestTimeout(ctx *fiber.Ctx) error
}
