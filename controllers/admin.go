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

func (adminC *AdminController) Reject(c *gin.Context) {
	oppID := c.Param("opp_id")
	reqBody := struct {
		RejectionReason string `json:"rejection_reason" bson:"rejection_reason" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := adminModel.Reject(oppID, reqBody.RejectionReason); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "success.")
}

func (adminC *AdminController) Undo(c *gin.Context) {
	oppID := c.Param("opp_id")
	if err := adminModel.Undo(oppID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "success.")
}

func (adminC *AdminController) Update(c *gin.Context) {
	adminID := c.Param("id")
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedAdmin, err := adminModel.Update(adminID, admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedAdmin)
}
