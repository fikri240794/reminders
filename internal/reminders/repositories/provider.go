package repositories

import (
	"github.com/fikri240794/reminders/datasources/reminders_sql_db/storages"
	"github.com/fikri240794/reminders/internal/reminders/persistences"
)

type ReminderRepository struct {
	reminderSQLDBStorage storages.IReminderSQLDBStorage
	reminderPersistence  persistences.IReminderPersistence
}

func NewReminderRepository(
	reminderSQLDBStorage storages.IReminderSQLDBStorage,
	reminderPersistence persistences.IReminderPersistence,
) *ReminderRepository {
	return &ReminderRepository{
		reminderSQLDBStorage: reminderSQLDBStorage,
		reminderPersistence:  reminderPersistence,
	}
}
