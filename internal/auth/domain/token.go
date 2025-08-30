package domain

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
