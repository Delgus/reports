package web

import (
	"net/http"
)

// Server - main server of our application
type Server struct {
	r1 *ReportHandler1
	r2 *ReportHandler2
}

// NewServer return new Server
func NewServer(r1 *ReportHandler1, r2 *ReportHandler2) *Server {
	return &Server{r1: r1, r2: r2}
}

// Serve - listen and serve requests
func (s *Server) Serve(addr string) error {
	http.HandleFunc("/r1/json", s.r1.JSON)
	http.HandleFunc("/r1/xlsx", s.r1.XLSX)
	http.HandleFunc("/r2/json", s.r2.JSON)
	http.HandleFunc("/r2/xlsx", s.r2.XLSX)

	return http.ListenAndServe(addr, nil)
}
