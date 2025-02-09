// internal/service/book.go
package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
)

// 2. BookService Struct
type BookService struct {
	repo *repository.BookRepository
}

// Constructor Function: NewBookService
func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

// GetBooks Method
func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetBookByID(id)
}

// 7. UpdateBook Method
func (s *BookService) UpdateBook(book *models.Book) error {
	return s.repo.UpdateBook(book)
}

// 8. DeleteBook Method
func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
