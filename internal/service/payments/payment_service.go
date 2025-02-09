package payments

import (
	"database/sql"
	"errors"
)

// PaymentService handles all payment-related operations.
type PaymentService struct {
	DB *sql.DB
}

// NewPaymentService initializes and returns a new PaymentService
func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{DB: db}
}

// OrderExists checks if an order exists
func (s *PaymentService) OrderExists(orderID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM orders WHERE id = $1)"
	err := s.DB.QueryRow(query, orderID).Scan(&exists)
	return exists, err
}

// IsPaymentAlreadyMade checks if a payment exists for an order
func (s *PaymentService) IsPaymentAlreadyMade(orderID int) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM payments WHERE order_id = $1"
	err := s.DB.QueryRow(query, orderID).Scan(&count)
	return count > 0, err //  Make sure it returns (bool, error)
}

// ProcessPayment inserts a new payment record
func (s *PaymentService) ProcessPayment(orderID int, userID int, amount float64) (int, error) {
	// Ensure order exists
	exists, err := s.OrderExists(orderID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, errors.New("order does not exist")
	}

	// Ensure payment is not duplicated
	alreadyPaid, err := s.IsPaymentAlreadyMade(orderID)
	if err != nil {
		return 0, err
	}
	if alreadyPaid {
		return 0, errors.New("payment already exists for this order")
	}

	// Insert payment
	query := "INSERT INTO payments (order_id, user_id, amount, status) VALUES ($1, $2, $3, 'completed') RETURNING id"
	var paymentID int
	err = s.DB.QueryRow(query, orderID, userID, amount).Scan(&paymentID)
	if err != nil {
		return 0, err
	}

	return paymentID, nil
}
