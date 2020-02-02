package tests

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/delgus/reports/config"
	"github.com/delgus/reports/internal/reporter1"
	"github.com/delgus/reports/internal/reporter2"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var (
	r1 *reporter1.Reporter
	r2 *reporter2.Reporter
)

func TestMain(m *testing.M) {
	flag.Parse()
	log.Print(`I work`)
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	var cfg config.Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Print(err)
		return 1
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDBName, cfg.PgPort)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
		return 1
	}
	r1 = reporter1.NewReporter(db)
	r2 = reporter2.NewReporter(db)
	return m.Run()
}
