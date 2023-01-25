package repositories

import (
	"context"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
)

func (r *ReminderRepository) countAllFromPersistence(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (int64, error) {
	var (
		count int64
		err   error
	)

	count, err = r.reminderPersistence.CountAllReminder(ctx, dto)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ReminderRepository) CountAll(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (int64, error) {
	var (
		count int64
		err   error
	)

	if dto == nil {
		return 0, errors.NewError(base_enums.BAD_REQUEST, "reminder pagination params dto is null")
	}

	count, err = r.countAllFromPersistence(ctx, dto)

	if err != nil {
		return 0, err
	}

	return count, nil
}
