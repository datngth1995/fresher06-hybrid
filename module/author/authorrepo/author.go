package authorrepo

import (
	"programming_fresher/common"
	"programming_fresher/token"
)

type LoginUserStorage interface {
	FindUserStorage(conditions map[string]interface{}) (*common.UserInfo, error)
}

type loginUserAuthorRepo struct {
	store LoginUserStorage
}

func NewLoginAuthorUserRepo(store LoginUserStorage) *loginUserAuthorRepo {
	return &loginUserAuthorRepo{store: store}
}

func (repo *loginUserAuthorRepo) Validate(payload *token.JWTPayLoad) (*common.UserInfo, error) {
	user, err := repo.store.FindUserStorage(map[string]interface{}{"id": payload.UserId})

	if err != nil {
		return nil, common.ErrorAuthentication(err)
	}

	if user.DeletedAt != nil {
		return nil, common.ErrorUserDeleted()
	}
	return user, nil
}
