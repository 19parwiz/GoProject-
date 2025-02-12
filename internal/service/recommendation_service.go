package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
)

type RecommendationService struct {
	Repo *repository.RecommendationRepository
}

// NewRecommendationService creates a new instance of RecommendationService
func NewRecommendationService(repo *repository.RecommendationRepository) *RecommendationService {
	return &RecommendationService{Repo: repo}
}

// GetRecommendations returns book recommendations based on a user's purchase history
func (s *RecommendationService) GetRecommendations(userID int) ([]models.Book, error) {
	recommendedBooks, err := s.Repo.GetRecommendedBooks(userID)
	if err != nil {
		return nil, err
	}
	return recommendedBooks, nil
}
