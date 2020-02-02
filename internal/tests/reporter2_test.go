package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReporter2JSON(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/json", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}
	w := httptest.NewRecorder()
	r2.JSON(w, req)
	if exp, got := http.StatusOK, w.Code; exp != got {
		t.Errorf("expected status code: %v, got status code: %v", exp, got)
	}
}
