package main

import (
	"fmt"
	"net/http"
	"programming_fresher/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect to database
	db, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}

	fmt.Println(db, err)

	//Running gin router
	r := gin.Default()

	v1 := r.Group("/fresher")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "This is pikachu"})
		})
	}

	r.Run()
}
