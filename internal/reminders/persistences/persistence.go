package persistences

import (
	"context"
	"database/sql"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

type IReminderPersistence interface {
	FindAllReminder(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) ([]*entities.ReminderEntity, error)
	CountAllReminder(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (int64, error)
	FindOneReminder(ctx context.Context, keyword string) (*entities.ReminderEntity, error)
	FindReminderById(ctx context.Context, id string) (*entities.ReminderEntity, error)
	CreateReminder(ctx context.Context, tx *sql.Tx, e *entities.ReminderEntity) error
	UpdateReminder(ctx context.Context, tx *sql.Tx, e *entities.ReminderEntity) error
}
