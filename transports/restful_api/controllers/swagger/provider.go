package swagger

import "github.com/gofiber/fiber/v2"

//	@title			Reminders RESTful API
//	@description	This is a Reminders RESTful API.

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
