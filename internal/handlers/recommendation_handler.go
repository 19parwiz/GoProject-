package handlers

import (
	"bookstore/internal/service"
	"encoding/json"
	"log" // Added missing log import
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
	vars := mux.Vars(r)
	userIDStr := vars["userID"]
	log.Println("Extracted userID from request:", userIDStr)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Error converting userID to integer:", err)
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	log.Println("Gathering the best recommendations tailored for userID", userID)

	recommendations, err := h.Service.GetRecommendations(userID)
	if err != nil {
		log.Printf("Error fetching recommendations: %v", err)
		http.Error(w, "Unable to retrieve recommendations at the moment", http.StatusInternalServerError)
		return
	}

	log.Println("Great! Recommendations have been successfully retrieved.:", recommendations)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
