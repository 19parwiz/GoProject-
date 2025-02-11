package handlers

import (
	"bookstore/internal/service/payments"
	"encoding/json"
	"net/http"
)

// PaymentHandler struct
type PaymentHandler struct {
	PaymentService *payments.PaymentService // Service to process payments
}

// NewPaymentHandler initializes a new PaymentHandler.
func NewPaymentHandler(paymentService *payments.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentService: paymentService}
}

// HandlePayment processes the payment for an order
func (h *PaymentHandler) HandlePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OrderID int     `json:"order_id"`
		UserID  int     `json:"user_id"`
		Amount  float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the payment using the provided details
	paymentID, err := h.PaymentService.ProcessPayment(req.OrderID, req.UserID, req.Amount)
	if err != nil {
		http.Error(w, "Payment failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Payment successful",
		"paymentID": paymentID,
	})
}
