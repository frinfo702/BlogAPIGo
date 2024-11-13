package services

import "database/sql"

type MyappService struct {
	db *sql.DB
}

// constructor
func NewAppService(db *sql.DB) *MyappService {
	return &MyappService{db: db}
}
