package favicons

import "github.com/gofiber/fiber/v2"

type IFaviconMiddleware interface {
	Favicon(ctx *fiber.Ctx) error
}
