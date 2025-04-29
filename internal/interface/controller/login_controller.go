package controller

// import (
// 	"api/internal/entity/request"
// 	"api/internal/usecase"
// 	"encoding/json"
// 	"net/http"

// 	"gorm.io/gorm"
// )

import (
	"api/config"
	"api/internal/entity"
	"api/internal/entity/request"
	"api/internal/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
    AuthUsecase *usecase.AuthUsecase
}

func NewAuthController(db *gorm.DB) *AuthController {
    return &AuthController{
        AuthUsecase: usecase.NewAuthUsecase(db),
    }
}

// Handler Login
// func LoginHandler(db *gorm.DB) http.HandlerFunc {
func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()

    var req request.LoginRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // Cari user di database
        var user entity.User
        if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        // Verifikasi password
        if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        // Buat custom claims
        expirationTime := time.Now().Add(24 * time.Hour)
        claims := &entity.CustomClaims{
            ID: user.ID,
            Username: user.Username,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
                IssuedAt:  time.Now().Unix(),
                Issuer:    "myapp",
                Subject:   fmt.Sprintf("%d", user.ID),
            },
        }

        // Buat token dengan HS256
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

        // Secret Key dari env
        secretKey := []byte("rahasia123") // Ganti dengan os.Getenv("JWT_SECRET") jika pakai .env

        // Tandatangani token
        tokenString, err := token.SignedString(secretKey)
        if err != nil {
            http.Error(w, "Failed to generate token", http.StatusInternalServerError)
            return
        }

        // Kirim token sebagai response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(entity.LoginResponse{
            Token: tokenString,
        })
}