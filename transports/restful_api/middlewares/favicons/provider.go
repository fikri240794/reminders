package favicons

import (
	"github.com/gofiber/fiber/v2"
	fiber_favicon "github.com/gofiber/fiber/v2/middleware/favicon"
)

type FaviconMiddleware struct {
	httpServer          *fiber.App
	fiberFaviconHandler func(ctx *fiber.Ctx) error
}

func NewFaviconMiddleware(
	httpServer *fiber.App,
) *FaviconMiddleware {
	var middleware *FaviconMiddleware = &FaviconMiddleware{
		httpServer:          httpServer,
		fiberFaviconHandler: fiber_favicon.New(),
	}

	middleware.Routing()

	return middleware
}
