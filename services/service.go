package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

// constructor
func NewAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
