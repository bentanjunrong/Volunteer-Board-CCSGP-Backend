package controllers

import (
	"context"
	"net/http"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oppModel.Create(ctx, opp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"opp_id": res.InsertedID})
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

func (oppC *OppController) Search(c *gin.Context) {
	params := c.Request.URL.Query()
	val, ok := params["query"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Did not include query param.",
		})
		return
	}
	matchedOpps, err := oppModel.Search(val[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"opportunities": matchedOpps})
}

func (oppC *OppController) GetOne(c *gin.Context) {
	params := c.Request.URL.Query()
	val, ok := params["id"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Did not include id param.",
		})
		return
	}
	matchedOpp, err := oppModel.GetOne(val[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"opportunity": matchedOpp})
}
