package requestids

import "github.com/gofiber/fiber/v2"

func (m *RequestIDMiddleware) RequestID(ctx *fiber.Ctx) error {
	return m.fiberRequestIDHandler(ctx)
}
