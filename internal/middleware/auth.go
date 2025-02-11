package middleware

// 1. Importing Necessary Packages
import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// 2. Defining the Secret Key
var jwtSecret = []byte("your_secret_key") //

// 3. Creating the AuthMiddleware Function
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем заголовок Authorization  4. Extracting the Authorization Header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		// 5. Checking the Header Format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		//6. Parsing and Validating the JWT
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNoCookie
			}
			return jwtSecret, nil
		})

		//7. Checking Token Validity
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		//  8. Allowing Access if Authentication Succeeds
		next.ServeHTTP(w, r)
	})
}
