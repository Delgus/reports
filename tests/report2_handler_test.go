package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/delgus/reports/internal/reports/report2"
	"github.com/delgus/reports/web"
)

func TestReporter2JSON(t *testing.T) {
	testJSON := `{"categories":[{"name":"Пиццы","products":[{"name":"4сыра","count":3,"cost_sum":"451.38","sell_sum":"1350.80"},{"name":"Мясное Плато","count":6,"cost_sum":"901.03","sell_sum":"2850.59"}],"count":9,"cost_sum":"1352.41","sell_sum":"4201.38"},{"name":"Супы","products":[{"name":"Борщ","count":3,"cost_sum":"90.99","sell_sum":"300.29"},{"name":"Харчо","count":3,"cost_sum":"60.51","sell_sum":"200.59"}],"count":6,"cost_sum":"151.50","sell_sum":"500.88"}],"count":15,"cost_sum":"1503.91","sell_sum":"4702.26"}`

	req, err := http.NewRequest(http.MethodGet, "/json", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	w := httptest.NewRecorder()

	reporter := web.NewReportHandler2(report2.NewService(GetDB()))
	reporter.JSON(w, req)

	if exp, got := http.StatusOK, w.Code; exp != got {
		t.Errorf("expected status code: %v, got status code: %v", exp, got)
	}

	answer := w.Body.String()

	if strings.TrimSpace(answer) != strings.TrimSpace(testJSON) {
		t.Errorf("unexpected response expect - %s got - %s", testJSON, answer)
	}
}

func TestReporter2XLSX(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/xlsx", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	w := httptest.NewRecorder()

	reporter := web.NewReportHandler2(report2.NewService(GetDB()))
	reporter.XLSX(w, req)

	if exp, got := http.StatusOK, w.Code; exp != got {
		t.Errorf("expected status code: %v, got status code: %v", exp, got)
	}
}
