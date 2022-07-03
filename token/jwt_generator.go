package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTgenerator struct {
	secretKey string
}

func NewJWTgenerator(s string) Generator {
	return &JWTgenerator{secretKey: s}
}

func (generator *JWTgenerator) CreateToken(userId int, role int, duration time.Duration) (string, error) {
	payload := NewJWTPayLoad(userId, role, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(generator.secretKey))
}

func (generator *JWTgenerator) VerifyToken(token string) (*JWTPayLoad, error) {

}
