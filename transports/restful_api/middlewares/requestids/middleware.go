package requestids

import "github.com/gofiber/fiber/v2"

type IRequestIDMiddleware interface {
	RequestID(ctx *fiber.Ctx) error
}
