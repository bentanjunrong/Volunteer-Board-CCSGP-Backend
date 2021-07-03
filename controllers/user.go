package controllers

import (
	"net/http"
	"os"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

var userModel = new(models.User)

func (userC *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedUser, err := userModel.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": savedUser})
}

func (userC *UserController) Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := userModel.Read(login.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot find user.",
		})
		return
	}

	userID := user["_id"].(string)
	userBody := user["_source"].(map[string]interface{})
	userPass := userBody["password"].(string)

	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(login.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password.",
		})
		return
	}

	jwtWrapper := utils.JwtWrapper{
		SecretKey:       os.Getenv("JWT_KEY"),
		Issuer:          os.Getenv("JWT_ISSUER"),
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(userID)
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
