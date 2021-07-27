package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type OppController struct{}

var oppModel = new(models.Opportunity)

func (oppC *OppController) GetAll(c *gin.Context) {
	allOpps, err := oppModel.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, allOpps)
}

func (oppC *OppController) GetAllApproved(c *gin.Context) {
	allOpps, err := oppModel.GetAllApproved()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, allOpps)
}

func (oppC *OppController) GetAllPending(c *gin.Context) {
	allOpps, err := oppModel.GetAllPending()
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
	c.JSON(http.StatusOK, matchedOpps)
}

func (oppC *OppController) GetOne(c *gin.Context) {
	id := c.Param("id")

	opp, err := oppModel.GetOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, opp)
}

func (oppC *OppController) CreateShift(c *gin.Context) {
	id := c.Param("id")
	var shift models.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := oppModel.CreateShift(id, shift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, "success.")
}

func (oppC *OppController) DeleteShift(c *gin.Context) {
	oppID := c.Param("id")
	shiftID := c.Param("shift_id")
	if err := oppModel.DeleteShift(oppID, shiftID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.String(http.StatusOK, "success.")
}

func (oppC *OppController) Update(c *gin.Context) { // TODO: only allow updates to certain fields!
	oppID := c.Param("id")
	var opp models.Opportunity
	if err := c.ShouldBindJSON(&opp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedOpp, err := oppModel.Update(oppID, opp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedOpp)
}
