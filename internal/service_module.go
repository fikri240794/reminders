package internal

import (
	"github.com/fikri240794/reminders/internal/reminders/services"
	"github.com/google/wire"
)

var ServiceModule wire.ProviderSet = wire.NewSet(
	services.NewReminderService,
	wire.Bind(new(services.IReminderService), new(*services.ReminderService)),
)
