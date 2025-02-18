package main

import (
	"bookstore/internal/handlers"
	"bookstore/internal/middleware"
	"bookstore/internal/repository"
	"bookstore/internal/service"
	"bookstore/internal/service/payments"
	"bookstore/pkg/config"
	"bookstore/pkg/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Configuring Database Connection
func main() {
	config.LoadConfig()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Initializing Repositories, Services, and Handlers
	//User Management

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	registrationHandler := handlers.NewRegistrationHandler(userService)
	loginHandler := handlers.NewLoginHandler(userService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	paymentService := payments.NewPaymentService(db) // Directly using db
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	// Book Management
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	mailHandler := handlers.NewMailHandler()

	// Defining API Routes
	r := mux.NewRouter()

	// Включаем CORS Middleware для всех маршрутов
	r.Use(middleware.CORSMiddleware)

	// Book API Routes (Protected by Authentication Middleware)
	r.Handle("/books", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.CreateBook))).Methods("POST")
	r.Handle("/books", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.GetBooks))).Methods("GET")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.GetBookByID))).Methods("GET")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.UpdateBook))).Methods("PUT")
	r.Handle("/books/{id}", middleware.AuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods("DELETE")

	// User Authentication & Email Routes
	r.HandleFunc("/send-email", mailHandler.SendEmail).Methods("POST")
	r.HandleFunc("/register", registrationHandler.Register).Methods("POST")
	r.HandleFunc("/login", loginHandler.Login).Methods("POST")
	r.HandleFunc("/logout", loginHandler.Logout).Methods("POST")

	r.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(orderHandler.CreateOrder))).Methods("POST")
	r.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(orderHandler.GetOrders))).Methods("GET")
	r.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(orderHandler.GetOrderByID))).Methods("GET")
	r.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(orderHandler.UpdateOrderStatus))).Methods("PUT")

	// Payments API route
	r.HandleFunc("/api/payment", paymentHandler.HandlePayment).Methods("POST")

	// Главная страница (добавляем явный маршрут для "/")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/home.html")
	}).Methods("GET")

	// Маршрут для получения данных пользователя
	r.Handle("/api/user", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := middleware.GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := userService.GetUserByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}))).Methods("GET")

	// Теперь раздаём статические файлы
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	//  Starting the HTTP Server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
