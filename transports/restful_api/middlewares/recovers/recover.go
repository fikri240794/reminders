package recovers

import "github.com/gofiber/fiber/v2"

func (m *RecoverMiddleware) Recover(ctx *fiber.Ctx) error {
	return m.fiberRecoverHandler(ctx)
}
