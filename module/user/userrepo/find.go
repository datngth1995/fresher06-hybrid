package userrepo

import (
	"programming_fresher/common"

	"golang.org/x/crypto/bcrypt"
)

type FindUser interface {
	findUserStorage(conditions map[string]interface{}) (*common.UserInfo, error)
}

type findUser struct {
	store FindUser
}

func NewFindUser(s FindUser) *findUser {
	return &findUser{store: s}
}

func (s *findUser) FindUser(userLogin *common.UserInfo) (*common.Account, error) {
	data, err := s.store.findUserStorage(map[string]interface{}{"email": userLogin.Email})

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password),
		[]byte(userLogin.Password)); err != nil {
		return nil, common.ErrorAuthentication(err)
	}

	//generate token

}
