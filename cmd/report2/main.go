package main

import (
	"net/http"

	report "github.com/delgus/reports/reports/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	reporter := report.NewReporter(db)
	http.HandleFunc("/json", reporter.JSON)
	http.HandleFunc("/xlsx", reporter.XLSX)
	if err := http.ListenAndServe(":8010", nil); err != nil {
		panic(err)
	}
}
