package userrepo

import (
	"programming_fresher/common"
)

type FindUser interface {
	FindUserStorage(conditions map[string]interface{}) (*common.UserInfo, error)
}

type findUser struct {
	store FindUser
}

func NewFindUser(s FindUser) *findUser {
	return &findUser{store: s}
}

func (s *findUser) FindUser(conditions map[string]interface{}) (*common.UserInfo, error) {
	data, err := s.store.FindUserStorage(conditions)

	if err != nil {
		return nil, err
	}

	return data, nil

}
