package configs

type RESTfulAPIServerConfig struct {
	Name                            string `mapstructure:"NAME"`
	Port                            int64  `mapstructure:"PORT"`
	Prefork                         bool   `mapstructure:"PREFORK"`
	PrintRoutes                     bool   `mapstructure:"PRINT_ROUTES"`
	RequestTimeoutInSec             int64  `mapstructure:"REQUEST_TIMEOUT_IN_SEC"`
	GracefullyShutdownDurationInSec int64  `mapstructure:"GRACEFULLY_SHUTDOWN_DURATION_IN_SEC"`
}

type RemindersSQLDBConnectionConfig struct {
	Host                 string `mapstructure:"HOST"`
	Port                 int64  `mapstructure:"PORT"`
	Username             string `mapstructure:"USERNAME"`
	Password             string `mapstructure:"PASSWORD"`
	Name                 string `mapstructure:"NAME"`
	MaxOpenConn          int64  `mapstructure:"MAX_OPEN_CONN"`
	ConnMaxLifetimeInSec int64  `mapstructure:"CONN_MAX_LIFE_TIME_IN_SEC"`
	MaxIdleConn          int64  `mapstructure:"MAX_IDLE_CONN"`
	ConnMaxIdleTimeInSec int64  `mapstructure:"CONN_MAX_IDLE_TIME_IN_SEC"`
}

type RemindersSQLDBConfig struct {
	Driver             string                         `mapstructure:"DRIVER"`
	MigrationSourceURL string                         `mapstructure:"MIGRATION_SOURCE_URL"`
	Master             RemindersSQLDBConnectionConfig `mapstructure:"MASTER"`
	Slave              RemindersSQLDBConnectionConfig `mapstructure:"SLAVE"`
}

type Config struct {
	Environment      string                  `mapstructure:"ENVIRONMENT"`
	RESTfulAPIServer *RESTfulAPIServerConfig `mapstructure:"RESTFUL_API_SERVER"`
	RemindersSQLDB   *RemindersSQLDBConfig   `mapstructure:"REMINDERS_SQL_DB"`
}
