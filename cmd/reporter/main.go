package main

import (
	"net/http"

	"github.com/delgus/reports/internal/reports/report1"
	"github.com/delgus/reports/internal/reports/report2"
	"github.com/delgus/reports/web"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

func main() {
	// configuration
	cfg, err := getConfig()
	if err != nil {
		logrus.Fatalf(`configuration error: %v`, err)
	}

	// create connections
	db := pg.Connect(&pg.Options{
		User:     cfg.PgUser,
		Password: cfg.PgPassword,
		Database: cfg.PgDBName,
		Addr:     cfg.PgAddr,
	})
	defer db.Close()

	// services
	reporter1 := report1.NewService(db)
	reporter2 := report2.NewService(db)

	// handlers
	reportHandler1 := web.NewReportHandler1(reporter1)
	reportHandler2 := web.NewReportHandler2(reporter2)

	// server
	server := web.NewServer(reportHandler1, reportHandler2)
	if err := server.Serve(cfg.AppPort); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err)
	}
}
