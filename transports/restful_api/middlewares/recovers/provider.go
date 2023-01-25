package recovers

import (
	"github.com/gofiber/fiber/v2"
	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"
)

type RecoverMiddleware struct {
	httpServer          *fiber.App
	fiberRecoverHandler func(ctx *fiber.Ctx) error
}

func NewRecoverMiddleware(
	httpServer *fiber.App,
) *RecoverMiddleware {
	var middleware *RecoverMiddleware = &RecoverMiddleware{
		httpServer:          httpServer,
		fiberRecoverHandler: fiber_recover.New(),
	}

	middleware.Routing()

	return middleware
}
