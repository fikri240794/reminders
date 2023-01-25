package persistences

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (p *ReminderPersistence) UpdateReminder(ctx context.Context, tx *sql.Tx, e *entities.ReminderEntity) error {
	var (
		query string
		args  []interface{}
		err   error
	)

	query, args, err = updateReminderQuery(p.cfg.RemindersSQLDB.Driver, e)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)

	if err != nil {
		return errors.NewError(base_enums.BAD_REQUEST, err.Error())
	}

	return nil
}
