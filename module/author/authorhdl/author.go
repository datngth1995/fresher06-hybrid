package authorhdl

import (
	"programming_fresher/common"
	"programming_fresher/token"
)

type LoginUserAuthorHdl interface {
	Validate(payload *token.JWTPayLoad) (*common.UserInfo, error)
}

type loginUserAuthorHdl struct {
	repo LoginUserAuthorHdl
}

func NewLoginUserAuthorHdl(repo LoginUserAuthorHdl) *loginUserAuthorHdl {
	return &loginUserAuthorHdl{repo: repo}
}

func (hdl *loginUserAuthorHdl) Validate(payload *token.JWTPayLoad) (*common.UserInfo, error) {
	user, err := hdl.repo.Validate(payload)

	if err != nil {
		return nil, err
	}

	return user, nil
}
