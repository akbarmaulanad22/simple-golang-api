package middleware

import (
	"api/internal/entity"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Key type untuk menyimpan data di context (menghindari clash)
// type contextKey string

// const UserKey contextKey = "user"

// AuthMiddleware adalah fungsi middleware untuk memverifikasi JWT
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        // Parse JWT
        // Parse JWT dengan custom claims
        token, err := jwt.ParseWithClaims(tokenString, &entity.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte("rahasia123"), nil // Ganti dengan env var
        })

        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }
        
        // if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        //     // Simpan username ke dalam context
        //     username, ok := claims["username"].(string)
        //     if !ok {
        //         http.Error(w, "Invalid token claims", http.StatusUnauthorized)
        //         return
        //     }

        //     ctx := context.WithValue(r.Context(), UserKey, username)
        //     next.ServeHTTP(w, r.WithContext(ctx))
        // } else {
        //     http.Error(w, "Invalid token", http.StatusUnauthorized)
        // }

        if claims, ok := token.Claims.(*entity.CustomClaims); ok && token.Valid {
            ctx := context.WithValue(r.Context(), "user", claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
        }
        
    })
}