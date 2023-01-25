package reminders

import (
	"github.com/fikri240794/reminders/internal/reminders/services"
	"github.com/gofiber/fiber/v2"
)

type ReminderController struct {
	httpServer      *fiber.App
	reminderService services.IReminderService
}

func NewReminderController(
	httpServer *fiber.App,
	reminderService services.IReminderService,
) *ReminderController {
	var controller *ReminderController = &ReminderController{
		httpServer:      httpServer,
		reminderService: reminderService,
	}

	controller.Routing()

	return controller
}
