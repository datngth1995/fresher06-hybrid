package main

import (
	"fmt"
	"programming_fresher/db"
	"programming_fresher/middleware"
	"programming_fresher/module/authen/authenhdl"
	"programming_fresher/module/user/userhdl"

	docs "programming_fresher/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const secretKey = "123456789"

func main() {
	//Connect to database
	db, err := db.ConnectToDB()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}

	//Running gin router
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/fresher"

	v1 := r.Group("/fresher")
	{
		v1.POST("/login", authenhdl.Login(db, secretKey))
		v1.GET("/me", middleware.RequiredAuth(db, secretKey), userhdl.GetProfile(db))
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
