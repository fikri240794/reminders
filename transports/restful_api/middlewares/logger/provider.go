package logger

import (
	"github.com/gofiber/fiber/v2"
)

type LoggerMiddleware struct {
	httpServer *fiber.App
}

func NewLoggerMiddleware(
	httpServer *fiber.App,
) *LoggerMiddleware {
	var middleware *LoggerMiddleware = &LoggerMiddleware{
		httpServer: httpServer,
	}

	middleware.Routing()

	return middleware
}
