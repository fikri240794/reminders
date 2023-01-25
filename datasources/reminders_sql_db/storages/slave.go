package storages

import "database/sql"

func (s *ReminderSQLDBStorage) Slave() *sql.DB {
	return s.slave
}
