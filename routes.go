package main

import (
	"log"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
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

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("ec2-13-212-235-180.ap-southeast-1.compute.amazonaws.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(router, &m))
}
