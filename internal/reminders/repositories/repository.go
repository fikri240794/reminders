package repositories

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

type IReminderRepository interface {
	FindAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) ([]*entities.ReminderEntity, error)
	CountAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (int64, error)
	FindOne(ctx context.Context, keyword string) (*entities.ReminderEntity, error)
	FindById(ctx context.Context, id string) (*entities.ReminderEntity, error)
	Create(ctx context.Context, e *entities.ReminderEntity) error
	Update(ctx context.Context, e *entities.ReminderEntity) error
	Delete(ctx context.Context, e *entities.ReminderEntity) error
}
