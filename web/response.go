package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

// Error is struct for displaing error
type Error struct {
	Message string `json:"message"`
}

func sendOK(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	sendJSON(w, response)
}

func sendJSON(w io.Writer, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.Error(err)
	}
}

func sendInternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	sendJSON(w, &Error{Message: "internal error"})
}

func sendXLSX(w http.ResponseWriter, r *http.Request, report *xlsx.File) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	err := report.Write(w)
	if err != nil {
		sendInternalError(w)
		logErrorWithRequest(err, r)
	}
}

func logErrorWithRequest(err error, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"request_uri":    r.RequestURI,
		"request_method": r.Method,
		"params":         r.Form,
	}).Error(err)
}
