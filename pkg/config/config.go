package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret string

// LoadConfig загружает переменные окружения из .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JwtSecret = os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in .env file")
	} else {
		log.Println("Loaded JWT Secret:", JwtSecret) // Отладочный вывод
	}
}

// GetDBConnectionString возвращает строку подключения к БД
func GetDBConnectionString() string {
	return os.Getenv("DATABASE_URL")
}
