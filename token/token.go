package token

import "time"

type JWTPayLoad struct {
	UserId    int       `json:"id"`
	Role      int       `json:"role"`
	IssuedAt  time.Time `json:"iat,omitempty"`
	ExpiredAt time.Time `json:"exp,omitempty"`
}

func NewJWTPayLoad(userId int, role int, duration time.Duration) *JWTPayLoad {
	return &JWTPayLoad{UserId: userId,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration)}
}
