package v1

import (
	"encoding/json"
	"log"
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
	report, err := r.GetJson()
	if err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Println(err)
	}
}

func (r *Reporter) XLSX(w http.ResponseWriter, req *http.Request) {

}
