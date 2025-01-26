// cmd/api/main.go
package main

import (
	"bookstore/internal/handlers"
	"bookstore/internal/middleware"
	"bookstore/internal/repository"
	"bookstore/internal/service"
	"bookstore/pkg/config"
	"bookstore/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	r := mux.NewRouter()
	r.Handle("/books", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.CreateBook))).Methods("POST")
	r.Handle("/books", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.GetBooks))).Methods("GET")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.GetBookByID))).Methods("GET")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.UpdateBook))).Methods("PUT")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
