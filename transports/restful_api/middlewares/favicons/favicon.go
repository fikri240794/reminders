package favicons

import "github.com/gofiber/fiber/v2"

func (m *FaviconMiddleware) Favicon(ctx *fiber.Ctx) error {
	return m.fiberFaviconHandler(ctx)
}
