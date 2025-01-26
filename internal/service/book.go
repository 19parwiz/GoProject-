// internal/service/book.go
package service

import (
	"bookstore/internal/models"
	"bookstore/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
