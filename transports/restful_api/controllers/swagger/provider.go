package swagger

import "github.com/gofiber/fiber/v2"

type SwaggerController struct {
	httpServer *fiber.App
}

func NewSwaggerController(httpServer *fiber.App) *SwaggerController {
	var controller *SwaggerController = &SwaggerController{
		httpServer: httpServer,
	}

	controller.Routing()

	return controller
}
