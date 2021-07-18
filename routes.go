package main

import (
	"os"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	users := router.Group("users")
	{
		userController := new(controllers.UserController)
		users.GET("/:id/opps", userController.GetOpps)
		users.PUT("/:id/apply/:opp_id", userController.ApplyOpp)
		users.PUT("/:id", userController.Update)
	}

	opps := router.Group("opps")
	{
		oppController := new(controllers.OppController)
		opps.GET("", oppController.GetAll)
		opps.GET("/pending", oppController.GetAllPending)
		opps.POST("", oppController.Create)
		opps.GET("/search", oppController.Search)
		opps.GET("/:id", oppController.GetOne)
		opps.POST("/:id/shifts", oppController.CreateShift)
		opps.DELETE("/:id/shifts/:shift_id", oppController.DeleteShift)
		opps.PUT("/:id", oppController.Update)
	}

	admins := router.Group("admins")
	{
		adminController := new(controllers.AdminController)
		admins.PUT("/approve/:opp_id", adminController.Approve)
		admins.PUT("/reject/:opp_id", adminController.Reject)
		admins.PUT("/:id", adminController.Update)
	}

	orgs := router.Group("orgs")
	{
		orgController := new(controllers.OrgController)
		orgs.PUT("/:id", orgController.Update)
	}

	// health check route for the LB
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "im alive!",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
