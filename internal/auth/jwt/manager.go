package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kitoyanok66/inote-auth/internal/auth/domain"
)

type JWTManager interface {
	Generate(userID, email string) (string, error)
	Verify(tokenStr string) (*domain.TokenClaims, error)
}

type jwtManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, duration time.Duration) JWTManager {
	return &jwtManager{secretKey: secretKey, tokenDuration: duration}
}

func (j *jwtManager) Generate(userID, email string) (string, error) {
	claims := &domain.TokenClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtManager) Verify(tokenStr string) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}
