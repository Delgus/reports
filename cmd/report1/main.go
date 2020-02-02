package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/delgus/reports/config"
	report "github.com/delgus/reports/internal/reporter1"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func main() {
	var cfg config.Configuration
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
