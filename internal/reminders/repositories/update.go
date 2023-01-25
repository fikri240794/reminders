package repositories

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (r *ReminderRepository) updateToPersistence(ctx context.Context, e *entities.ReminderEntity) error {
	var (
		tx          *sql.Tx
		rollbackErr error
		err         error
	)

	tx, err = r.reminderSQLDBStorage.BeginTx(ctx)

	if err != nil {
		return err
	}

	err = r.reminderPersistence.UpdateReminder(ctx, tx, e)

	if err != nil {
		rollbackErr = tx.Rollback()

		if rollbackErr != nil {
			return errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
		}

		return err
	}

	err = tx.Commit()

	if err != nil {
		return errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return nil
}

func (r *ReminderRepository) Update(ctx context.Context, e *entities.ReminderEntity) error {
	var err error

	if e == nil {
		return errors.NewError(base_enums.BAD_REQUEST, "reminder entity is null")
	}

	err = r.updateToPersistence(ctx, e)

	if err != nil {
		return err
	}

	return nil
}
