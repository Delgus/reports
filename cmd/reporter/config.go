package main

import "github.com/kelseyhightower/envconfig"

// Configuration is config for app
type configuration struct {
	PgAddr     string `envconfig:"PG_ADDR" default:"localhost:5432"`
	PgUser     string `envconfig:"PG_USER" default:"postgres"`
	PgPassword string `envconfig:"PG_PASSWORD" default:"123456"`
	PgDBName   string `envconfig:"PG_DBNAME" default:"postgres"`
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
