package controllers

import (
	"github.com/fikri240794/reminders/transports/restful_api/controllers/reminders"
	"github.com/fikri240794/reminders/transports/restful_api/controllers/swagger"
	"github.com/google/wire"
)

var ControllerModule wire.ProviderSet = wire.NewSet(
	reminders.NewReminderController,
	wire.Bind(new(reminders.IReminderController), new(*reminders.ReminderController)),
	swagger.NewSwaggerController,
	wire.Bind(new(swagger.ISwaggerController), new(*swagger.SwaggerController)),
	NewController,
)
