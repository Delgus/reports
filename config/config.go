package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

// Configuration is config for app
type Configuration struct {
	PgHost     string `envconfig:"PG_HOST" default:"localhost"`
	PgUser     string `envconfig:"PG_USER" default:"postgres"`
	PgPassword string `envconfig:"PG_PASSWORD" default:"123456"`
	PgDBName   string `envconfig:"PG_DBNAME" default:"postgres"`
	PgPort     int    `envconfig:"PG_PORT" default:"5432"`
	AppPort    int    `envconfig:"APP_PORT" default:"80"`
}

// GetConfig get configuration from environments
func GetConfig() (*Configuration, error) {
	var cfg Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetDBConnection get db connection from configuration
func GetDBConnection(cfg *Configuration) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDBName, cfg.PgPort)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
