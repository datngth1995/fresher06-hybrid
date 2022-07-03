package userstorage

import "gorm.io/gorm"

type userMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) *userMySQL {
	return &userMySQL{db: db}
}
