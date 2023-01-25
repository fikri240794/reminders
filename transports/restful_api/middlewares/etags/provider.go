package etags

import (
	"github.com/gofiber/fiber/v2"
	fiber_etag "github.com/gofiber/fiber/v2/middleware/etag"
)

type ETagMiddleware struct {
	httpServer       *fiber.App
	fiberETagHandler func(ctx *fiber.Ctx) error
}

func NewETagMiddleware(
	httpServer *fiber.App,
) *ETagMiddleware {
	var middleware *ETagMiddleware = &ETagMiddleware{
		httpServer:       httpServer,
		fiberETagHandler: fiber_etag.New(),
	}

	middleware.Routing()

	return middleware
}
