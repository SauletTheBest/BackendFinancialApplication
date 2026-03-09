package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct{
	secret string
}

func NewService(secret string) *Service {
	return &Service{secret: secret}
}

func (s *Service) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(s.secret))
}