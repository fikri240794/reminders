package internal

import (
	"github.com/fikri240794/reminders/internal/reminders/persistences"
	"github.com/google/wire"
)

var PersistenceModule wire.ProviderSet = wire.NewSet(
	persistences.NewReminderPersistence,
	wire.Bind(new(persistences.IReminderPersistence), new(*persistences.ReminderPersistence)),
)
