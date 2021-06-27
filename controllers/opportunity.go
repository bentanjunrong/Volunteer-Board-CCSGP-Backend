package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type OppController struct{}

var oppModel = new(models.Opportunity)

func (oppC *OppController) Create(c *gin.Context) {
	var opp models.Opportunity
	if err := c.ShouldBindJSON(&opp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := oppModel.Create(opp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"opportunity": opp})
}

func (oppC *OppController) GetAll(c *gin.Context) {
	allOpps, err := oppModel.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"opportunities": allOpps})
}
