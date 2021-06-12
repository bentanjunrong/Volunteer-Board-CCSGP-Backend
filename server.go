package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "why am i doing this at 4 in the morning?",
		})
	})

	server.Run("localhost:8080")
}
