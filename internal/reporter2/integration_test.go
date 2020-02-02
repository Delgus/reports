package reporter2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/delgus/reports/config"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var testJSON2 = `{"categories":[{"name":"Пиццы","products":[{"name":"4сыра","count":3,"cost_sum":"451.38","sell_sum":"1350.80"},{"name":"Мясное Плато","count":6,"cost_sum":"901.03","sell_sum":"2850.59"}],"count":9,"cost_sum":"1352.41","sell_sum":"4201.38"},{"name":"Супы","products":[{"name":"Борщ","count":3,"cost_sum":"90.99","sell_sum":"300.29"},{"name":"Харчо","count":3,"cost_sum":"60.51","sell_sum":"200.59"}],"count":6,"cost_sum":"151.50","sell_sum":"500.88"}],"count":15,"cost_sum":"1503.91","sell_sum":"4702.26"}`

func TestReporter2JSON(t *testing.T) {
	// test database
	var cfg config.Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		t.Error(err)
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		cfg.TestPgHost, cfg.TestPgUser, cfg.TestPgPassword, cfg.TestPgDBName, cfg.TestPgPort)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// reporter
	reporter := NewReporter(db)

	req, err := http.NewRequest(http.MethodGet, "/json", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}
	w := httptest.NewRecorder()
	reporter.JSON(w, req)
	if exp, got := http.StatusOK, w.Code; exp != got {
		t.Errorf("expected status code: %v, got status code: %v", exp, got)
	}
	answer := w.Body.String()
	if strings.TrimSpace(answer) != strings.TrimSpace(testJSON2) {
		t.Errorf("unexpected response expect - %s got - %s", testJSON2, answer)
	}
}
