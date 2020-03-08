package web

import (
	"net/http"

	"github.com/delgus/reports/internal/reports/report2"
)

// ReportHandler2 - report handler
type ReportHandler2 struct {
	service *report2.Service
}

// NewReportHandler2 return new service Reporter
func NewReportHandler2(s *report2.Service) *ReportHandler2 {
	return &ReportHandler2{service: s}
}

// JSON return report in JSON
func (r *ReportHandler2) JSON(w http.ResponseWriter, req *http.Request) {
	report, err := r.service.GetJSON()
	if err != nil {
		sendInternalError(w)
		logErrorWithRequest(err, req)

		return
	}

	sendOK(w, report)
}

// XLSX return report in xlsx
func (r *ReportHandler2) XLSX(w http.ResponseWriter, req *http.Request) {
	report, err := r.service.GetXLSX()
	if err != nil {
		sendInternalError(w)
		logErrorWithRequest(err, req)

		return
	}

	sendXLSX(w, req, report)
}
