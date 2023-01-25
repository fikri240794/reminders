package etags

import "github.com/gofiber/fiber/v2"

func (m *ETagMiddleware) ETag(ctx *fiber.Ctx) error {
	return m.fiberETagHandler(ctx)
}
