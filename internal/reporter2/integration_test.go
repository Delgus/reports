package reporter2

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/delgus/reports/config"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testJSON = `{"categories":[{"name":"Пиццы","products":[{"name":"4сыра","count":3,"cost_sum":"451.38","sell_sum":"1350.80"},{"name":"Мясное Плато","count":6,"cost_sum":"901.03","sell_sum":"2850.59"}],"count":9,"cost_sum":"1352.41","sell_sum":"4201.38"},{"name":"Супы","products":[{"name":"Борщ","count":3,"cost_sum":"90.99","sell_sum":"300.29"},{"name":"Харчо","count":3,"cost_sum":"60.51","sell_sum":"200.59"}],"count":6,"cost_sum":"151.50","sell_sum":"500.88"}],"count":15,"cost_sum":"1503.91","sell_sum":"4702.26"}`

func TestReporter2JSON(t *testing.T) {
	cfg, err := config.GetConfig()
	if err != nil {
		t.Errorf(`configuration error: %v`, err)
	}

	db, err := config.GetDBConnection(cfg)
	if err != nil {
		t.Errorf(`db open connection error: %v`, err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf(`db close connection error: %v`, err)
		}
	}()

	// reporter
	reporter := NewReporter(db)

	req, err := http.NewRequest(http.MethodGet, "/json", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}
	w := httptest.NewRecorder()
	reporter.JSON(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testJSON, strings.TrimSpace(w.Body.String()))
}
