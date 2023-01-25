package services

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
)

type IReminderService interface {
	FindAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (*dtos.ReminderPaginationResultDTO, error)
	FindOne(ctx context.Context, keyword string) (*dtos.ReminderDTO, error)
	FindById(ctx context.Context, id string) (*dtos.ReminderDTO, error)
	Create(ctx context.Context, dto *dtos.CreateReminderDTO) (*dtos.ReminderDTO, error)
	Update(ctx context.Context, dto *dtos.UpdateReminderDTO) (*dtos.ReminderDTO, error)
	Delete(ctx context.Context, id string) error
}
