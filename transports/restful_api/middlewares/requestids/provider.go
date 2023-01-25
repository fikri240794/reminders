package requestids

import (
	"github.com/gofiber/fiber/v2"
	fiber_requestid "github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type RequestIDMiddleware struct {
	httpServer            *fiber.App
	fiberRequestIDHandler func(ctx *fiber.Ctx) error
}

func NewRequestIDMiddleware(
	httpServer *fiber.App,
) *RequestIDMiddleware {
	var middleware *RequestIDMiddleware = &RequestIDMiddleware{
		httpServer: httpServer,
		fiberRequestIDHandler: fiber_requestid.New(
			fiber_requestid.Config{
				Header: fiber.HeaderXRequestID,
				Generator: func() string {
					return uuid.NewString()
				},
				ContextKey: "RequestID",
			},
		),
	}

	middleware.Routing()

	return middleware
}
