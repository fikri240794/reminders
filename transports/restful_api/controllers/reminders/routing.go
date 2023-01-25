package reminders

import "github.com/gofiber/fiber/v2"

func (c *ReminderController) Routing() {
	var rg fiber.Router = c.httpServer.Group("/reminders")

	rg.Post("/", c.Create)
	rg.Get("/", c.FindAll)
	rg.Get("/find", c.FindOne)
	rg.Get("/:id", c.FindById)
	rg.Put("/:id", c.Update)
	rg.Delete("/:id", c.Delete)
}
