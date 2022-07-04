package middleware

import (
	"errors"
	"net/http"
	"programming_fresher/common"
	"programming_fresher/module/author/authorhdl"
	"programming_fresher/module/author/authorrepo"
	"programming_fresher/module/user/userstorage"
	"programming_fresher/token"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequiredAuth(db *gorm.DB, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithError(http.StatusUnauthorized, common.ErrorTokenInvalid())
			c.JSON(http.StatusBadRequest, common.ErrorAuthentication(errors.New("Invalid token")))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2) {
			c.AbortWithError(http.StatusUnauthorized, common.ErrorTokenInvalid())
			c.JSON(http.StatusBadRequest, common.ErrorAuthentication(errors.New("Invalid token")))
			return
		}

		tokenGenerator := token.NewJWTgenerator(secretKey)
		payload, err := tokenGenerator.VerifyToken(parts[1])

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, common.ErrorTokenInvalid())
			c.JSON(http.StatusBadRequest, common.ErrorAuthentication(err))
			return
		}

		store := userstorage.NewUserMySQL(db)
		repo := authorrepo.NewLoginAuthorUserRepo(store)
		hdl := authorhdl.NewLoginUserAuthorHdl(repo)

		user, err := hdl.Validate(payload)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, common.ErrorUserDeleted())
			c.JSON(http.StatusBadRequest, common.ErrorAuthentication(err))
			return
		}
		c.Set("data", user)
		c.Next()
	}
}
