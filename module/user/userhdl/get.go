package userhdl

import (
	"errors"
	"net/http"
	"programming_fresher/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("data")

		if !ok {
			c.JSON(http.StatusBadRequest, common.ErrorInternal(errors.New("Something went wrong")))
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
