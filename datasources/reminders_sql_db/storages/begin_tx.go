package storages

import (
	"context"
	"database/sql"

	"github.com/fikri240794/golib/errors"
	base_enums "github.com/fikri240794/golib/models/enums"
)

func (s *ReminderSQLDBStorage) BeginTx(ctx context.Context) (*sql.Tx, error) {
	var (
		tx  *sql.Tx
		err error
	)

	tx, err = s.master.BeginTx(ctx, nil)

	if err != nil {
		return nil, errors.NewError(base_enums.INTERNAL_SERVER_ERROR, err.Error())
	}

	return tx, nil
}
