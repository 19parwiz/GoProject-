package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
	"log"
)

// RecommendationService struct
type RecommendationService struct {
	Repo *repository.RecommendationRepository
}

// ✅ Constructor function
func NewRecommendationService(repo *repository.RecommendationRepository) *RecommendationService {
	return &RecommendationService{Repo: repo}
}

// Fetch recommendations for a user
func (s *RecommendationService) GetRecommendations(userID int) ([]models.Book, error) {
	log.Println("Retrieving recommendations for the user..:", userID) // ✅ Log user ID

	recommendedBooks, err := s.Repo.GetRecommendedBooks(userID)
	if err != nil {
		log.Println("Error in GetRecommendations:", err) // ✅ Log DB error
		return nil, err
	}

	log.Println("Your recommended books are ready!:", recommendedBooks) // ✅ Log response
	return recommendedBooks, nil
}
