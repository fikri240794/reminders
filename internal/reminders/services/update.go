package services

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (i *ReminderService) Update(ctx context.Context, dto *dtos.UpdateReminderDTO) (*dtos.ReminderDTO, error) {
	var (
		e      *entities.ReminderEntity
		result *dtos.ReminderDTO
		err    error
	)

	err = dto.Validate()

	if err != nil {
		return nil, err
	}

	e, err = i.reminderRepository.FindById(ctx, dto.ID)

	if err != nil {
		return nil, err
	}

	e = dto.MapToReminderEntity(e)
	err = i.reminderRepository.Update(ctx, e)

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
