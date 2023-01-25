package internal

import (
	"github.com/fikri240794/reminders/internal/reminders/repositories"
	"github.com/google/wire"
)

var RepositoryModule wire.ProviderSet = wire.NewSet(
	repositories.NewReminderRepository,
	wire.Bind(new(repositories.IReminderRepository), new(*repositories.ReminderRepository)),
)
