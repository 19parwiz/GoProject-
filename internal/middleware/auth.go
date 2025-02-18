package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"bookstore/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, загружен ли `JWT_SECRET`
		if config.JwtSecret == "" {
			fmt.Println(" JWT_SECRET is empty in authMiddleware! Check config.go")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Получаем заголовок Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println(" Missing Authorization header")
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Разделяем строку `Bearer <token>`
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println(" Invalid Authorization header format:", authHeader)
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		fmt.Println(" Received Token:", tokenString)
		fmt.Println(" JWT Secret in authMiddleware:", config.JwtSecret)

		// Разбираем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println(" Token validation error:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Получаем user_id из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println(" Invalid token claims")
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println(" user_id not found in token claims")
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)
		fmt.Println(" Extracted user_id from token:", userID)

		// Передаём user_id в контекст (если нужно использовать в хендлерах)
		r.Header.Set("X-User-ID", fmt.Sprintf("%d", userID))

		next.ServeHTTP(w, r)
	})
}

// Функция для извлечения user_id из токена (используется в API)
func GetUserIDFromToken(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, fmt.Errorf("invalid Authorization header format")
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token")
	}

	return int(userIDFloat), nil
}
