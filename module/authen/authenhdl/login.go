package authenhdl

import (
	"net/http"
	"programming_fresher/common"
	"programming_fresher/module/authen/authenrepo"
	"programming_fresher/module/user/userstorage"
	"programming_fresher/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Login godoc
// @Summary ping example
// @Router /login [POST]
// @Schemes
// @Description Authentication login: email and password
// @Param string query string
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
func Login(db *gorm.DB, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData common.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorInternal(err))
			return
		}

		tokenGenerator := token.NewJWTgenerator(secretKey)
		store := userstorage.NewUserMySQL(db)
		repo := authenrepo.NewLoginAuthenUserRepo(store, tokenGenerator)

		account, err := repo.LoginAuthenUserRepo(&loginUserData)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrorAuthentication(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"access_token": account.AccessToken})
	}
}
