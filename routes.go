package main

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

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

	opp := router.Group("opp")
	{
		oppController := new(controllers.OppController)
		opp.GET("", oppController.GetAll)
		opp.POST("", oppController.Create)
		opp.GET("/search", oppController.Search)
	}

	router.Run(":8080") // TODO: put this in an env file
}
