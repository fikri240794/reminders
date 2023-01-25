package persistences

import (
	"github.com/fikri240794/reminders/configs"
	"github.com/fikri240794/reminders/datasources/reminders_sql_db/storages"
)

type ReminderPersistence struct {
	cfg                  *configs.Config
	reminderSQLDBStorage storages.IReminderSQLDBStorage
}

func NewReminderPersistence(
	cfg *configs.Config,
	ReminderSQLDBStorage storages.IReminderSQLDBStorage,
) *ReminderPersistence {
	return &ReminderPersistence{
		cfg:                  cfg,
		reminderSQLDBStorage: ReminderSQLDBStorage,
	}
}
