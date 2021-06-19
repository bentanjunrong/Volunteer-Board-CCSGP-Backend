package main

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := router.Group("auth")
	{
		authController := new(controllers.AuthController)
		auth.POST("", authController.Register)
		auth.POST("login", authController.Login)
	}

	user := router.Group("user")
	{
		userController := new(controllers.UserController)
		user.GET("", userController.Retrieve)
	}

	router.Run("localhost:8080") // TODO: put this in an env file
}
