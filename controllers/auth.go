package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

type AuthController struct{}

var authModel = new(models.User)

func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return (*token).Raw, nil
}

func (j *JwtWrapper) ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		return errors.New("Error when parsing claims.")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("JWT expired.")
	}

	return nil
}

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
