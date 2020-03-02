package report1

import "github.com/go-pg/pg"

// Service - service for build report
type Service struct {
	store *pg.DB
}

// NewService return new service Service
func NewService(store *pg.DB) *Service {
	return &Service{store: store}
}
