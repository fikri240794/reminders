package request_timeouts

import (
	"github.com/fikri240794/reminders/configs"
	"github.com/gofiber/fiber/v2"
)

type RequestTimeoutMiddleware struct {
	cfg        *configs.Config
	httpServer *fiber.App
}

func NewRequestTimeoutMiddleware(
	cfg *configs.Config,
	httpServer *fiber.App,
) *RequestTimeoutMiddleware {
	var middleware *RequestTimeoutMiddleware = &RequestTimeoutMiddleware{
		cfg:        cfg,
		httpServer: httpServer,
	}

	middleware.Routing()

	return middleware
}
