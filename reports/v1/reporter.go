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

//JSON вернет отчет в формате JSON
func (r *Reporter) JSON(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	report, err := r.getJson()
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

//XSLX вернет отчет в формате XLSX
func (r *Reporter) XLSX(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	report, err := r.getXLSX()
	if err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
	}

	if err := report.Write(w); err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
	}
}
