package handlers

import (
	"bookstore/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RecommendationHandler struct {
	Service *service.RecommendationService
}

// NewRecommendationHandler creates a new handler for recommendations
func NewRecommendationHandler(service *service.RecommendationService) *RecommendationHandler {
	return &RecommendationHandler{Service: service}
}

// GetRecommendations handles the request to fetch book recommendations
func (h *RecommendationHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL parameters
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch recommendations
	recommendedBooks, err := h.Service.GetRecommendations(userID)
	if err != nil {
		http.Error(w, "Failed to get recommendations", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendedBooks)
}
