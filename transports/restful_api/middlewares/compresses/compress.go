package compresses

import "github.com/gofiber/fiber/v2"

func (m *CompressMiddleware) Compress(ctx *fiber.Ctx) error {
	return m.fiberCompressHandler(ctx)
}
