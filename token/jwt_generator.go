package token

import (
	"programming_fresher/common"
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

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		/*_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, common.ErrorTokenInvalid()
		}*/
		return []byte(generator.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayLoad{}, keyFunc)

	if err != nil {
		return nil, common.ErrorTokenInvalid()
	}

	payload, ok := jwtToken.Claims.(*JWTPayLoad)

	if !ok {
		return nil, common.ErrorTokenInvalid()
	}

	return payload, nil
}
