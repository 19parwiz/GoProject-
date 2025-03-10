// internal/repository/book.go      Explanation of BookRepository (internal/repository/book.go)
package repository

import (
	"bookstore/internal/models"
	"database/sql"
)

// 2. BookRepository Struct
type BookRepository struct {
	db *sql.DB
}

// 3. Constructor Function: NewBookRepository
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

// 4. CreateBook Method
func (r *BookRepository) CreateBook(book *models.Book) error {
	query := `INSERT INTO books (title, author, published_date, price) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, book.Title, book.Author, book.PublishedDate, book.Price).Scan(&book.ID)
}

// 5. GetBooks Method
func (r *BookRepository) GetBooks() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, published_date, price FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Price); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// 6. GetBookByID Method
func (r *BookRepository) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	query := `SELECT id, title, author, published_date, price FROM books WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.Price)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// 7. UpdateBook Method
func (r *BookRepository) UpdateBook(book *models.Book) error {
	query := `UPDATE books SET title = $1, author = $2, published_date = $3, price = $4 WHERE id = $5`
	_, err := r.db.Exec(query, book.Title, book.Author, book.PublishedDate, book.Price, book.ID)
	return err
}

// 8. DeleteBook Method
func (r *BookRepository) DeleteBook(id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
