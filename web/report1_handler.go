package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/delgus/reports/internal/reports/report1"
)

// ReportHandler1 - report handler
type ReportHandler1 struct {
	service *report1.Service
}

// NewReportHandler1 return new service Reporter
func NewReportHandler1(s *report1.Service) *ReportHandler1 {
	return &ReportHandler1{service: s}
}

// JSON return report in JSON
func (r *ReportHandler1) JSON(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	report, err := r.service.GetJSON()
	if err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Println(err)
	}
}

// XLSX return report in xlsx
func (r *ReportHandler1) XLSX(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	report, err := r.service.GetXLSX()
	if err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
		return
	}

	if err := report.Write(w); err != nil {
		log.Println(err)
		if err := json.NewEncoder(w).Encode(&Error{Message: "Ooops!"}); err != nil {
			log.Println(err)
		}
	}
}
