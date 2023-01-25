package storages

import "database/sql"

func (s *ReminderSQLDBStorage) Master() *sql.DB {
	return s.master
}
