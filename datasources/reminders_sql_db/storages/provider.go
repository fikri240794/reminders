package storages

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fikri240794/reminders/configs"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type ReminderSQLDBStorage struct {
	master *sql.DB
	slave  *sql.DB
}

func NewReminderSQLDBStorage(cfg *configs.Config) *ReminderSQLDBStorage {
	var (
		storage   *ReminderSQLDBStorage
		dsnf      string
		masterDSN string
		slaveDSN  string
		errGroup  *errgroup.Group
		err       error
	)

	storage = &ReminderSQLDBStorage{}

	if cfg.RemindersSQLDB.Driver == "postgres" {
		dsnf = "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
	} else {
		dsnf = "%s:%s@tcp(%s:%d)/%s"
	}

	masterDSN = fmt.Sprintf(
		dsnf,
		cfg.RemindersSQLDB.Master.Username,
		cfg.RemindersSQLDB.Master.Password,
		cfg.RemindersSQLDB.Master.Host,
		cfg.RemindersSQLDB.Master.Port,
		cfg.RemindersSQLDB.Master.Name,
	)
	slaveDSN = fmt.Sprintf(
		dsnf,
		cfg.RemindersSQLDB.Slave.Username,
		cfg.RemindersSQLDB.Slave.Password,
		cfg.RemindersSQLDB.Slave.Host,
		cfg.RemindersSQLDB.Slave.Port,
		cfg.RemindersSQLDB.Slave.Name,
	)
	errGroup = &errgroup.Group{}
	errGroup.Go(func() error {
		var errRoutine error

		storage.master, err = sql.Open(
			cfg.RemindersSQLDB.Driver,
			masterDSN,
		)

		if errRoutine != nil {
			return errRoutine
		}

		storage.master.SetMaxOpenConns(int(cfg.RemindersSQLDB.Master.MaxOpenConn))
		storage.master.SetConnMaxLifetime(time.Duration(cfg.RemindersSQLDB.Master.ConnMaxLifetimeInSec) * time.Second)
		storage.master.SetMaxIdleConns(int(cfg.RemindersSQLDB.Master.MaxIdleConn))
		storage.master.SetConnMaxIdleTime(time.Duration(cfg.RemindersSQLDB.Master.ConnMaxIdleTimeInSec) * time.Second)

		return nil
	})
	errGroup.Go(func() error {
		var errRoutine error

		storage.slave, err = sql.Open(
			cfg.RemindersSQLDB.Driver,
			slaveDSN,
		)

		if errRoutine != nil {
			return errRoutine
		}

		storage.slave.SetMaxOpenConns(int(cfg.RemindersSQLDB.Slave.MaxOpenConn))
		storage.slave.SetConnMaxLifetime(time.Duration(cfg.RemindersSQLDB.Slave.ConnMaxLifetimeInSec) * time.Second)
		storage.slave.SetMaxIdleConns(int(cfg.RemindersSQLDB.Slave.MaxIdleConn))
		storage.slave.SetConnMaxIdleTime(time.Duration(cfg.RemindersSQLDB.Slave.ConnMaxIdleTimeInSec) * time.Second)

		return nil
	})
	err = errGroup.Wait()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	return storage
}
