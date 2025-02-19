package repository

import (
	"bookstore/internal/models"
	"database/sql"
	"log"
)

// RecommendationRepository struct
type RecommendationRepository struct {
	DB *sql.DB
}

// ✅ Constructor function
func NewRecommendationRepository(db *sql.DB) *RecommendationRepository {
	return &RecommendationRepository{DB: db}
}

// ✅ GetRecommendedBooks function
func (r *RecommendationRepository) GetRecommendedBooks(userID int) ([]models.Book, error) {
	log.Printf("Fetching recommended books for userID: %d", userID) // Debug log

	var books []models.Book
	query := `
    SELECT b.id, b.title, b.author, b.genre, b.price
    FROM books b
    JOIN orders o ON b.id = o.book_id
    WHERE o.user_id = $1
    GROUP BY b.id, b.title, b.author, b.genre, b.price
    ORDER BY COUNT(o.book_id) DESC
    LIMIT 5;
    `

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		log.Printf("❌ Error executing query: %s | userID: %d | Error: %v", query, userID, err)
		return nil, err
	}
	defer rows.Close() // ✅ Move this inside the correct scope

	// ✅ Fetch books from result set
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Price); err != nil {
			log.Println(" Error scanning recommended books:", err)
			return nil, err
		}
		log.Printf("✅ Book retrieved successfully: %+v", book) // Debug log
		books = append(books, book)
	}

	// ✅ Log when no recommendations are found
	if len(books) == 0 {
		log.Println(" No recommendations found for userID:", userID)
	}

	return books, nil
}
