package handlers

import (
	"bookstore/internal/models"
	"bookstore/internal/service"
	"encoding/json"
	"net/http"
)

type RegistrationHandler struct {
	UserService *service.UserService
}

func NewRegistrationHandler(userService *service.UserService) *RegistrationHandler {
	return &RegistrationHandler{UserService: userService}
}

func (h *RegistrationHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.UserService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}
