package services

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (i *ReminderService) FindOne(ctx context.Context, keyword string) (*dtos.ReminderDTO, error) {
	var (
		e      *entities.ReminderEntity
		result *dtos.ReminderDTO
		err    error
	)

	e, err = i.reminderRepository.FindOne(ctx, keyword)

	if err != nil {
		return nil, err
	}

	result = dtos.NewReminderDTO().
		MapFromReminderEntity(e)

	return result, nil
}
