package authenrepo

import (
	"programming_fresher/common"
	"programming_fresher/token"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserStorage interface {
	FindUserStorage(conditions map[string]interface{}) (*common.UserInfo, error)
}

type loginUserAuthenRepo struct {
	store          LoginUserStorage
	tokenGenerator token.Generator
}

func NewLoginAuthenUserRepo(store LoginUserStorage, tokenGenerator token.Generator) *loginUserAuthenRepo {
	return &loginUserAuthenRepo{store: store, tokenGenerator: tokenGenerator}
}

func (repo *loginUserAuthenRepo) LoginAuthenUserRepo(loginUserData *common.UserLogin) (*common.Account, error) {
	user, err := repo.store.FindUserStorage(map[string]interface{}{"email": loginUserData.Email})

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(loginUserData.Password)); err != nil {
		return nil, common.ErrorAuthentication(err)
	}

	accessToke, err := repo.tokenGenerator.CreateToken(user.Id, user.RoleId, token.Duration)

	if err != nil {
		return nil, common.ErrorInternal(err)
	}

	account := common.NewAccount(accessToke)

	return account, nil
}
