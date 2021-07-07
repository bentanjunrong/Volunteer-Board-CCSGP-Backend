package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (userC *UserController) GetOpps(c *gin.Context) {
	userID := c.Param("id")
	opps, err := userModel.GetOpps(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, opps)
}
