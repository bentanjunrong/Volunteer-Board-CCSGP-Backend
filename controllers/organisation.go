package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type OrgController struct{}

var orgModel = new(models.Organisation)

// TODO: refactor into a common function
func (orgC *OrgController) Register(c *gin.Context) {
	var org models.Organisation
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedOrg, err := orgModel.Create(org)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"org": savedOrg})
}

func (orgC *OrgController) Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	org, err := orgModel.Read(login.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find organisation.",
		})
		return
	}

	orgID := org["_id"].(string)
	orgBody := org["_source"].(map[string]interface{})
	orgPass := orgBody["password"].(string)

	if err := bcrypt.CompareHashAndPassword([]byte(orgPass), []byte(login.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password.",
		})
		return
	}

	jwtWrapper := utils.JwtWrapper{
		SecretKey:       "mySecretKey", // TODO: pls put in env
		Issuer:          "VolunteeryAuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Error signing token.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signedToken,
	})
}
