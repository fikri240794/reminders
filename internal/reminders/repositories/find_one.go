package repositories

import (
	"context"

	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (r *ReminderRepository) findOneFromPersistence(ctx context.Context, keyword string) (*entities.ReminderEntity, error) {
	var (
		e   *entities.ReminderEntity
		err error
	)

	e, err = r.reminderPersistence.FindOneReminder(ctx, keyword)

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *ReminderRepository) FindOne(ctx context.Context, keyword string) (*entities.ReminderEntity, error) {
	var (
		e   *entities.ReminderEntity
		err error
	)

	e, err = r.findOneFromPersistence(ctx, keyword)

	if err != nil {
		return nil, err
	}

	return e, nil
}
