package repositories

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (r *ReminderRepository) findByIdFromPersistence(ctx context.Context, id string) (*entities.ReminderEntity, error) {
	var (
		e   *entities.ReminderEntity
		err error
	)

	e, err = r.reminderPersistence.FindReminderById(ctx, id)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *ReminderRepository) FindById(ctx context.Context, id string) (*entities.ReminderEntity, error) {
	var (
		e   *entities.ReminderEntity
		err error
	)

	e, err = r.findByIdFromPersistence(ctx, id)

	if err != nil {
		return nil, err
	}

	return e, nil
}
