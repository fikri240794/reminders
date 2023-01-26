package database_migrations

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/fikri240794/reminders/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var (
	cfgPath         string
	migrationTarget string
	cfg             *configs.Config
	db              *sql.DB
	driver          database.Driver
	dbmigrate       *migrate.Migrate
	Cmd             *cobra.Command
)

func getVersionFromMigrationTarget() (uint, error) {
	var (
		splittedFileName []string
		version          uint64
		err              error
	)

	splittedFileName = strings.Split(migrationTarget, "_")

	if len(splittedFileName) == 0 {
		err = fmt.Errorf("%s file name is invalid", migrationTarget)

		return 0, err
	}

	version, err = strconv.ParseUint(splittedFileName[0], 10, 64)

	if err != nil {
		return 0, err
	}

	return uint(version), nil
}

func runMigration() error {
	var (
		migrationTargetVersion uint
		err                    error
	)

	if migrationTarget == "" {
		err = dbmigrate.Up()
	} else {
		migrationTargetVersion, err = getVersionFromMigrationTarget()

		if err != nil {
			return err
		}

		err = dbmigrate.Migrate(migrationTargetVersion)
	}

	if err != nil {
		return err
	}

	return nil
}

func preRun(cmd *cobra.Command, args []string) error {
	var (
		dsnf string
		dsn  string
		err  error
	)

	configs.SetConfigPath(cfgPath)
	cfg = configs.Read()

	if cfg.RemindersSQLDB.Driver == "postgres" {
		dsnf = "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
	} else {
		dsnf = "%s:%s@tcp(%s:%d)/%s"
	}

	dsn = fmt.Sprintf(
		dsnf,
		cfg.RemindersSQLDB.Master.Username,
		cfg.RemindersSQLDB.Master.Password,
		cfg.RemindersSQLDB.Master.Host,
		cfg.RemindersSQLDB.Master.Port,
		cfg.RemindersSQLDB.Master.Name,
	)

	db, err = sql.Open(cfg.RemindersSQLDB.Driver, dsn)

	if err != nil {
		return err
	}

	if cfg.RemindersSQLDB.Driver == "postgres" {
		driver, err = postgres.WithInstance(db, &postgres.Config{})
	} else {
		driver, err = mysql.WithInstance(db, &mysql.Config{})
	}

	if err != nil {
		return err
	}

	dbmigrate, err = migrate.NewWithDatabaseInstance(cfg.RemindersSQLDB.MigrationSourceURL, cfg.RemindersSQLDB.Driver, driver)

	if err != nil {
		return err
	}

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	return runMigration()
}

func postRun(cmd *cobra.Command, args []string) error {
	return db.Close()
}

func init() {
	Cmd = &cobra.Command{
		Use:      "database-migration",
		PreRunE:  preRun,
		RunE:     run,
		PostRunE: postRun,
	}

	Cmd.PersistentFlags().StringVar(
		&cfgPath,
		"config-path",
		".env",
		"config file path",
	)

	Cmd.PersistentFlags().StringVar(
		&migrationTarget,
		"target",
		"",
		"migration target version by migration file name, if not set, migration will run to latest version",
	)
}
