package config

// Configuration is config for app
type Configuration struct {
	PgHost     string `envconfig:"PG_HOST"`
	PgUser     string `envconfig:"PG_USER"`
	PgPassword string `envconfig:"PG_PASSWORD"`
	PgDBName   string `envconfig:"PG_DBNAME"`
	PgPort     int    `envconfig:"PG_PORT"`
	AppPort    int    `envconfig:"APP_PORT"`
}
