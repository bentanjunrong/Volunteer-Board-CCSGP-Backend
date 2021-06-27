package controllers

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/db"
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string
	Password string
}

type AuthController struct{}

var authModel = new(models.User)

func (a *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedUser, err := authModel.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": savedUser})
}

func (a *AuthController) Login(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := db.GetOne("users", map[string]string{
		"email": login.Email,
	})
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
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
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
