package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type OrgController struct{}

var orgModel = new(models.Organisation)

func (orgC *OrgController) GetOpps(c *gin.Context) {
	orgID := c.Param("id")
	opps, err := orgModel.GetOpps(orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, opps)
}

func (orgC *OrgController) Update(c *gin.Context) {
	orgID := c.Param("id")
	var org models.Organisation
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedOrg, err := orgModel.Update(orgID, org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedOrg)
}

func (orgC *OrgController) GetOne(c *gin.Context) {
	id := c.Param("id")

	user, err := orgModel.GetOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
