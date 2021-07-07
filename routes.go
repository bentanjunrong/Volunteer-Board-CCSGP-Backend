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

	opp := router.Group("opps")
	{
		oppController := new(controllers.OppController)
		opp.GET("", oppController.GetAll)
		opp.POST("", oppController.Create)
		opp.GET("/search", oppController.Search)
		opp.GET("/:id", oppController.GetOne)
		opp.POST("/:id/shifts", oppController.CreateShift)
		opp.DELETE("/:id/shifts/:shift_id", oppController.DeleteShift)
	}

	// health check route for the LB
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "im alive!",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
}
