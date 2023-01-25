package recovers

import "github.com/gofiber/fiber/v2"

type IRecoverMiddleware interface {
	Recover(ctx *fiber.Ctx) error
}
