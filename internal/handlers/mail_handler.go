package handlers

import (
	"bookstore/internal/utils"
	"encoding/json"
	"net/http"
)

type MailHandler struct{}

func NewMailHandler() *MailHandler {
	return &MailHandler{}
}

type SendMailRequest struct {
	To             string `json:"to"`
	Subject        string `json:"subject"`
	Body           string `json:"body"`
	AttachmentPath string `json:"attachment_path,omitempty"`
}

// SendEmail handles email sending requests
func (h *MailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	var req SendMailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := utils.SendMail(req.To, req.Subject, req.Body, req.AttachmentPath)
	if err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
}
