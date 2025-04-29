package entity

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}