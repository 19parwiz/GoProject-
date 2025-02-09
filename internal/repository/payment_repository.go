package repository

import "database/sql"

// PaymentRepository handles database operations for payments
type PaymentRepository struct {
	DB *sql.DB
}

// NewPaymentRepository creates a new repository instance
func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}
