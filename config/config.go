package config

// Configuration is config for app
type Configuration struct {
	PgHost         string `envconfig:"PG_HOST"`
	PgUser         string `envconfig:"PG_USER"`
	PgPassword     string `envconfig:"PG_PASSWORD"`
	PgDBName       string `envconfig:"PG_DBNAME"`
	PgPort         int    `envconfig:"PG_PORT"`
	TestPgHost     string `envconfig:"TEST_PG_HOST" default:"localhost"`
	TestPgUser     string `envconfig:"TEST_PG_USER" default:"postgres"`
	TestPgPassword string `envconfig:"TEST_PG_PASSWORD" default:"123456"`
	TestPgDBName   string `envconfig:"TEST_PG_DBNAME" default:"postgres"`
	TestPgPort     int    `envconfig:"TEST_PG_PORT" default:"5432"`
	AppPort        int    `envconfig:"APP_PORT"`
}
