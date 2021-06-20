package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authModel = new(models.User)

func (a *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := authModel.Create(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
