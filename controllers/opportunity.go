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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // can move these to models!
	defer cancel()
	oppID, err := oppModel.Create(ctx, opp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"opp_id": oppID})
}

func (oppC *OppController) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	allOpps, err := oppModel.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, allOpps)
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
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	opp, err := oppModel.GetOne(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, opp)
}
