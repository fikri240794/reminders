package persistences

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
)

func (p *ReminderPersistence) CountAllReminder(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) (int64, error) {
	var (
		query string
		args  []interface{}
		row   *sql.Row
		count int64
		err   error
	)

	query, args, err = countAllReminderQuery(p.cfg.RemindersSQLDB.Driver, dto)

	if err != nil {
		return 0, err
	}

	row = p.reminderSQLDBStorage.Slave().
		QueryRowContext(ctx, query, args...)
	err = row.Scan(&count)

	if err != nil {
		return 0, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	if count == 0 {
		return 0, errors.NewError(base_enums.NOT_FOUND, "notes not found")
	}

	return count, nil
}
