package main

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	user := router.Group("user")
	{
		userController := new(controllers.UserController)
		user.POST("", userController.Register)
		user.POST("login", userController.Login)
	}

	org := router.Group("org")
	{
		orgController := new(controllers.OrgController)
		org.POST("", orgController.Register)
		org.POST("login", orgController.Login)
	}

	router.Run("localhost:8080") // TODO: put this in an env file
}
