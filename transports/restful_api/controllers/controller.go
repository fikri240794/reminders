package controllers

import (
	"github.com/fikri240794/reminders/transports/restful_api/controllers/reminders"
	"github.com/fikri240794/reminders/transports/restful_api/controllers/swagger"
)

type Controller struct {
	swaggerController  swagger.ISwaggerController
	reminderController reminders.IReminderController
}

func NewController(
	swaggerController swagger.ISwaggerController,
	reminderController reminders.IReminderController,
) *Controller {
	return &Controller{
		swaggerController:  swaggerController,
		reminderController: reminderController,
	}
}
