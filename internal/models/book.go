// internal/models/book.go
package models

//  Defining the Book Struct
type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	PublishedDate string  `json:"published_date"`
	Price         float64 `json:"price"`
}
