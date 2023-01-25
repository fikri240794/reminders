package swagger

import (
	"github.com/gofiber/fiber/v2"
	fiber_swagger "github.com/gofiber/swagger"
)

func (c *SwaggerController) Index(ctx *fiber.Ctx) error {
	return fiber_swagger.HandlerDefault(ctx)
}
