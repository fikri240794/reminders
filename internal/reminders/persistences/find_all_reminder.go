package persistences

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
	"github.com/fikri240794/reminders/internal/reminders/models/dtos"
	"github.com/fikri240794/reminders/internal/reminders/models/entities"
)

func (p *ReminderPersistence) FindAllReminder(ctx context.Context, dto *dtos.ReminderPaginationParamsDTO) ([]*entities.ReminderEntity, error) {
	var (
		query       string
		args        []interface{}
		rows        *sql.Rows
		entityDatas []*entities.ReminderEntity
		err         error
	)

	query, args, err = findAllReminderQuery(p.cfg.RemindersSQLDB.Driver, dto)

	if err != nil {
		return nil, err
	}

	rows, err = p.reminderSQLDBStorage.Slave().
		QueryContext(ctx, query, args...)

	if err != nil {
		return nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	defer rows.Close()

	entityDatas = []*entities.ReminderEntity{}

	for rows.Next() {
		var e *entities.ReminderEntity = entities.NewReminderEntity()

		err = rows.Scan(
			&e.ID,
			&e.Note,
			&e.CreatedAt,
			&e.UpdatedAt,
			&e.IsDeleted,
		)

		if err != nil {
			return nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
		}

		entityDatas = append(entityDatas, e)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	if len(entityDatas) == 0 {
		return nil, errors.NewError(base_enums.NOT_FOUND, "notes not found")
	}

	return entityDatas, nil
}
