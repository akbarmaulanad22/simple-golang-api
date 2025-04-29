package usecase

import (
	"api/internal/repository"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUsecase struct {
    // Bisa inject repository lain di sini
    LoginRepository *repository.LoginRepository
}

func NewAuthUsecase(db *gorm.DB) *AuthUsecase {
    return &AuthUsecase{
        LoginRepository: repository.NewLoginRepository(db),
    }
}

type LoginResponse struct {
    Token string `json:"token"`
}

func (s *AuthUsecase) Login(username, password string) (*LoginResponse, error) {
    login, err := s.LoginRepository.FindByUsername(username)
    if err != nil {
        return nil, errors.New("invalid username or password")
    }

    // Verifikasi password
    if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid username or password")
    }

    // Buat JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": login.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    secretKey := []byte("mysecretpassword") // Ambil dari env sebenarnya
    tokenString, _ := token.SignedString(secretKey)

    return &LoginResponse{
        Token: tokenString,
    }, nil
}