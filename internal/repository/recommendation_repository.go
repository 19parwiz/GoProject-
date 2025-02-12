package repository

import (
	"bookstore/internal/models"
	"database/sql"
	"log"
)

type RecommendationRepository struct {
	DB *sql.DB
}

// NewRecommendationRepository creates a new instance of RecommendationRepository
func NewRecommendationRepository(db *sql.DB) *RecommendationRepository {
	return &RecommendationRepository{DB: db}
}

// GetUserPurchasedBooks fetches book IDs that the user has purchased
func (r *RecommendationRepository) GetUserPurchasedBooks(userID int) ([]int, error) {
	query := `SELECT book_id FROM orders WHERE user_id = ?`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookIDs []int
	for rows.Next() {
		var bookID int
		if err := rows.Scan(&bookID); err != nil {
			return nil, err
		}
		bookIDs = append(bookIDs, bookID)
	}

	return bookIDs, nil
}

// GetRecommendedBooks finds books bought by other users who purchased similar books
func (r *RecommendationRepository) GetRecommendedBooks(userID int) ([]models.Book, error) {
	query := `
		SELECT DISTINCT b.id, b.title, b.author, b.genre
	FROM orders o
	JOIN orders o2 ON o.book_id = o2.book_id AND o.user_id != o2.user_id
	JOIN books b ON o2.book_id = b.id
	WHERE o.user_id = 8
	LIMIT 5;

	`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		log.Println("Error fetching recommendations:", err)
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
