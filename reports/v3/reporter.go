package v3

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Reporter struct {
	store *sqlx.DB
}

func NewReporter(store *sqlx.DB) *Reporter {
	return &Reporter{store: store}
}

func (r *Reporter) JSON(w http.ResponseWriter, req *http.Request) {

}

func (r *Reporter) XLSX(w http.ResponseWriter, req *http.Request) {

}
