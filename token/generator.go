package token

import "time"

const Duration = 24 * 30 * time.Hour

type Generator interface {
	CreateToken(userId int, role int, duration time.Duration) (string, error)
	VerifyToken(token string) (*JWTPayLoad, error)
}
