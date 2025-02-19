package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"bookstore/pkg/config"

	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware is a middleware that checks and validates JWT tokens
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if `JWT_SECRET` is set
		if config.JwtSecret == "" {
			fmt.Println("JWT_SECRET is empty in AuthMiddleware! Check config.go")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("Missing Authorization header")
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Split the header value `Bearer <token>`
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Println("Invalid Authorization header format:", authHeader)
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Extract the token
		tokenString := parts[1]
		fmt.Println("Received Token:", tokenString)
		fmt.Println("JWT Secret in AuthMiddleware:", config.JwtSecret)

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("Token validation error:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user_id from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Invalid token claims")
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println("user_id not found in token claims")
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)
		fmt.Println("Extracted user_id from token:", userID)

		// Pass user_id in the request header (for use in handlers)
		r.Header.Set("X-User-ID", fmt.Sprintf("%d", userID))

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromToken extracts user_id from the JWT token (used in API)
func GetUserIDFromToken(r *http.Request) (int, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, fmt.Errorf("missing Authorization header")
	}

	// Split the header value `Bearer <token>`
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, fmt.Errorf("invalid Authorization header format")
	}

	// Extract the token
	tokenString := parts[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	// Extract user_id from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("user_id not found in token")
	}

	// Return user_id as an integer
	return int(userIDFloat), nil
}
