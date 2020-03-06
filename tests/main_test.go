package tests

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.BuildAndRun(
		"test-postgres",
		"./Dockerfile",
		[]string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=123456",
		},
	)
	if err != nil {
		logrus.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		source := fmt.Sprintf("host=localhost user=postgres password=123456 dbname=postgres sslmode=disable port=%s",
			resource.GetPort("5432/tcp"))
		db, err = sqlx.Open("pgx", source)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		logrus.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		logrus.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func GetDB() *sqlx.DB {
	return db
}
