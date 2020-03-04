package report1

import "github.com/jmoiron/sqlx"

// Service - service for build report
type Service struct {
	store *sqlx.DB
}

// NewService return new service Service
func NewService(store *sqlx.DB) *Service {
	return &Service{store: store}
}
