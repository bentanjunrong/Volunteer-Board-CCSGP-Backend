package controllers

import (
	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var authModel = new(models.User)

func (a *AuthController) Register(ctx *gin.Context) {}

func (a *AuthController) Login(ctx *gin.Context) {}
