package web

import (
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
	report, err := r.service.GetJSON()
	if err != nil {
		sendInternalError(w)
		logErrorWithRequest(err, req)

		return
	}

	sendOK(w, report)
}

// XLSX return report in xlsx
func (r *ReportHandler1) XLSX(w http.ResponseWriter, req *http.Request) {
	report, err := r.service.GetXLSX()
	if err != nil {
		sendInternalError(w)
		logErrorWithRequest(err, req)

		return
	}

	sendXLSX(w, req, report)
}
