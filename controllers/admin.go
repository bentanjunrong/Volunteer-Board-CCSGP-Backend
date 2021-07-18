package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

var adminModel = new(models.Admin)

func (adminC *AdminController) Approve(c *gin.Context) {
	oppID := c.Param("opp_id")
	if err := adminModel.Approve(oppID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "success.")
}
