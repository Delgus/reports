package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.BuildAndRun(
		"test-postgres",
		"./Dockerfile",
		[]string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=123456",
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	log.Println(resource.GetPort("5432/tcp"))

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
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
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func GetDB() *sqlx.DB {
	return db
}
