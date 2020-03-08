package main

import "github.com/kelseyhightower/envconfig"

type configuration struct {
	PgPort     int    `envconfig:"PG_PORT" default:"5432"`
	PgHost     string `envconfig:"PG_HOST" default:"localhost"`
	PgUser     string `envconfig:"PG_USER" default:"postgres"`
	PgPassword string `envconfig:"PG_PASSWORD" default:"123456"`
	PgDBName   string `envconfig:"PG_DBNAME" default:"postgres"`
	AppHost    string `envconfig:"APP_HOST" default:""`
	AppPort    int    `envconfig:"APP_PORT" default:"80"`
}

func getConfig() (*configuration, error) {
	var cfg configuration

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
