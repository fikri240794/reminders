package repositories

import (
	"context"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (r *ReminderRepository) Delete(ctx context.Context, e *entities.ReminderEntity) error {
	var err error

	if e == nil {
		return errors.NewError(base_enums.BAD_REQUEST, "reminder entity is null")
	}

	e = e.MarkAsDeleted()
	err = r.updateToPersistence(ctx, e)

	if err != nil {
		return err
	}

	return nil
}
