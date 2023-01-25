package persistences

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (p *ReminderPersistence) FindReminderById(ctx context.Context, id string) (*entities.ReminderEntity, error) {
	var (
		query string
		args  []interface{}
		row   *sql.Row
		e     *entities.ReminderEntity
		err   error
	)

	query, args, err = findReminderByIdQuery(p.cfg.RemindersSQLDB.Driver, id)

	if err != nil {
		return nil, err
	}

	row = p.reminderSQLDBStorage.Slave().
		QueryRowContext(ctx, query, args...)
	e = entities.NewReminderEntity()
	err = row.Scan(
		&e.ID,
		&e.Note,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.IsDeleted,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewError(base_enums.NOT_FOUND, "note not found")
		}

		return nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return e, nil
}
