package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc { // after middleware run handlerfunc
	return func(w http.ResponseWriter, r *http.Request) { // middleware return a func, output func
		fmt.Println("Entered in middleware....")
		authHeader := r.Header.Get("Authorization") // Get authorization value form browser URL
		fmt.Println("Authorization Header :", authHeader)

		if authHeader == "" { // authHeader is missing
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ") // remove bearer
		fmt.Println("Token :", tokenString)

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // jwt.Parse(check token)
			// in login token create, during verfiy with same key
			return []byte("my-secret-key"), nil
		})

		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		fmt.Println("JWT Verified Successfully")
		next(w, r) // end the middleware work , next controller
	}
}

/*
3 case tested
1. no header // unauthorized
2. Bearer abc123 // Invalid token
3. Copy token. Authorization: Bearer eyJhbGc... // success
*/
