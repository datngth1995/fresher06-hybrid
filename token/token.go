package token

import (
	"errors"
	"programming_fresher/common"
	"time"
)

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

func (payload *JWTPayLoad) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return common.ErrorTokenExpried(errors.New("token has expired"))
	}
	return nil
}
