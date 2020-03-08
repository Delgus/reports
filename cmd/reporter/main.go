package main

import (
	"fmt"
	"net/http"

	"github.com/delgus/reports/internal/reports/report1"
	"github.com/delgus/reports/internal/reports/report2"
	"github.com/delgus/reports/web"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	// configuration
	cfg, err := getConfig()
	if err != nil {
		logrus.Fatalf(`configuration error: %v`, err)
	}

	// create connections
	db, err := newDBConnection(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Error(err)
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
	addr := fmt.Sprintf(`%s:%d`, cfg.AppHost, cfg.AppPort)

	if err := server.Serve(addr); err != nil && err != http.ErrServerClosed {
		logrus.Fatal(err)
	}
}

func newDBConnection(cfg *configuration) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDBName, cfg.PgPort)

	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
