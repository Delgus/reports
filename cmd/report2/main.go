package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/delgus/reports/config"
	report "github.com/delgus/reports/internal/reporter2"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf(`configuration error: %v`, err)
	}

	db, err := config.GetDBConnection(cfg)
	if err != nil {
		log.Fatalf(`db open connection error: %v`, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf(`db close connection error: %v`, err)
		}
	}()

	reporter := report.NewReporter(db)
	http.HandleFunc("/json", reporter.JSON)
	http.HandleFunc("/xlsx", reporter.XLSX)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.AppPort), nil); err != nil {
		log.Println(err)
	}
}
