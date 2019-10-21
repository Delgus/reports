package main

import (
	"net/http"

	report "github.com/delgus/reports/internal/reporter2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=postgres user=postgres password=123456 dbname=postgres sslmode=disable port=5432"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	reporter := report.NewReporter(db)
	http.HandleFunc("/json", reporter.JSON)
	http.HandleFunc("/xlsx", reporter.XLSX)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
