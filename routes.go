package main

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	user := router.Group("user")
	{
		userController := new(controllers.UserController)
		user.GET("", userController.Retrieve)
	}

	return router
}
