package corses

import "github.com/gofiber/fiber/v2"

func (m *CORSMiddleware) CORS(ctx *fiber.Ctx) error {
	return m.fiberCORSHandler(ctx)
}
