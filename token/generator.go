package token

import "time"

type Generator interface {
	CreateToken(userId int, role int, duration time.Duration) (string, error)
	VerifyToken(token string) (*JWTPayLoad, error)
}
