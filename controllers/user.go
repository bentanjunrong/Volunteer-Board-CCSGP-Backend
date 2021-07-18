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

func (userC *UserController) ApplyOpp(c *gin.Context) {
	userID := c.Param("id")
	oppID := c.Param("opp_id")
	reqBody := struct {
		ShiftIDs []string `json:"shift_ids" bson:"shift_ids" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := userModel.ApplyOpp(userID, oppID, reqBody.ShiftIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, "success.")
}

func (userC *UserController) Update(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser, err := userModel.Update(userID, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
