package common

import (
	"go/token"
	"time"
)

type UserInfo struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Email     string     `json:"email" gorm:"column:email;"`
	Name      string     `json:"name" gorm:"colum:name;"`
	Phone     string     `json:"phone" gorm:"column:phone;"`
	RoleId    string     `json:"role_id" gorm:"column:role_id;"`
	Password  string     `json:"password" gorm:"column:password;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}

func (UserInfo) TableName() string {
	return "users"
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

type Account struct {
	AccessToken  *token.Token `json:"access_token"`
	RefreshToken *token.Token `json:"refresh_token"`
}

func NewAccount(atok, rtok *token.Token) *Account {
	return &Account{AccessToken: atok, RefreshToken: rtok}
}
