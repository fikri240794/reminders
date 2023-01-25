package corses

import (
	"github.com/gofiber/fiber/v2"
	fiber_cors "github.com/gofiber/fiber/v2/middleware/cors"
)

type CORSMiddleware struct {
	httpServer       *fiber.App
	fiberCORSHandler func(ctx *fiber.Ctx) error
}

func NewCORSMiddleware(
	httpServer *fiber.App,
) *CORSMiddleware {
	var middleware *CORSMiddleware = &CORSMiddleware{
		httpServer:       httpServer,
		fiberCORSHandler: fiber_cors.New(),
	}

	middleware.Routing()

	return middleware
}
