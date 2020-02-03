package main

import (
	"log"

	"github.com/delgus/reports/config"
	"github.com/delgus/reports/internal/reports/report1"
	"github.com/delgus/reports/internal/reports/report2"
	"github.com/delgus/reports/web"
	_ "github.com/lib/pq"
)

func main() {
	// configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf(`configuration error: %v`, err)
	}

	// create connections
	db, err := config.GetDBConnection(cfg)
	if err != nil {
		log.Fatalf(`db open connection error: %v`, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf(`db close connection error: %v`, err)
		}
	}()

	// services
	reporter1 := report1.NewService(db)
	reporter2 := report2.NewService(db)

	// handlers
	reportHandler1 := web.NewReportHandler1(reporter1)
	reportHandler2 := web.NewReportHandler2(reporter2)

	// server
	server := web.NewServer(reportHandler1, reportHandler2)
	if err := server.Serve(cfg.AppPort); err != nil {
		log.Println(err)
	}
}
