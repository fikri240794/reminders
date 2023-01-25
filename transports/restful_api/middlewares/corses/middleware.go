package corses

import "github.com/gofiber/fiber/v2"

type ICORSMiddleware interface {
	CORS(ctx *fiber.Ctx) error
}
