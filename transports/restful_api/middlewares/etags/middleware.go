package etags

import "github.com/gofiber/fiber/v2"

type IETagMiddleware interface {
	ETag(ctx *fiber.Ctx) error
}
