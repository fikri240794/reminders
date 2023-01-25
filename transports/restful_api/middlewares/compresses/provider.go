package compresses

import (
	"github.com/gofiber/fiber/v2"
	fiber_compress "github.com/gofiber/fiber/v2/middleware/compress"
)

type CompressMiddleware struct {
	httpServer           *fiber.App
	fiberCompressHandler func(ctx *fiber.Ctx) error
}

func NewCompressMiddleware(
	httpServer *fiber.App,
) *CompressMiddleware {
	var middleware *CompressMiddleware = &CompressMiddleware{
		httpServer: httpServer,
		fiberCompressHandler: fiber_compress.New(
			fiber_compress.Config{
				Level: fiber_compress.LevelBestSpeed,
			},
		),
	}

	middleware.Routing()

	return middleware
}
