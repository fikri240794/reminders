package datasources

import (
	"github.com/fikri240794/reminders/datasources/reminders_sql_db/storages"
	"github.com/google/wire"
)

var DatasourceModule wire.ProviderSet = wire.NewSet(
	storages.NewReminderSQLDBStorage,
	wire.Bind(new(storages.IReminderSQLDBStorage), new(*storages.ReminderSQLDBStorage)),
)
