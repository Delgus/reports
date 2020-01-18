package main

import (
	"fmt"
	"net/http"
	"strconv"

	report "github.com/delgus/reports/internal/reporter2"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type config struct {
	PgHost     string `envconfig:"PG_HOST"`
	PgUser     string `envconfig:"PG_USER"`
	PgPassword string `envconfig:"PG_PASSWORD"`
	PgDBName   string `envconfig:"PG_DBNAME"`
	PgPort     int    `envconfig:"PG_PORT"`
	AppPort    int    `envconfig:"APP_PORT"`
}

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDBName, cfg.PgPort)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	reporter := report.NewReporter(db)
	http.HandleFunc("/json", reporter.JSON)
	http.HandleFunc("/xlsx", reporter.XLSX)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.AppPort), nil); err != nil {
		panic(err)
	}
}
