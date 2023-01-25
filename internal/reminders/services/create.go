package services

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (i *ReminderService) Create(ctx context.Context, dto *dtos.CreateReminderDTO) (*dtos.ReminderDTO, error) {
	var (
		e      *entities.ReminderEntity
		result *dtos.ReminderDTO
		err    error
	)

	err = dto.Validate()

	if err != nil {
		return nil, err
	}

	e = dto.MapToReminderEntity()
	err = i.reminderRepository.Create(ctx, e)

	if err != nil {
		return nil, err
	}

	e, err = i.reminderRepository.FindById(ctx, e.ID)

	if err != nil {
		return nil, err
	}

	result = dtos.NewReminderDTO().
		MapFromReminderEntity(e)

	return result, nil
}
