package middlewares

import (
	"net/http"
	"strings"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/utils"
	"github.com/gin-gonic/gin"
)

// expects "Authorization" header for any requests that require auth with value "Bearer <JWT Token>".
func Auth(c *gin.Context) {
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	extractedToken := strings.Split(clientToken, "Bearer ")
	if len(extractedToken) <= 1 {
		c.JSON(http.StatusBadRequest, "Incorrect Format of Authorization Token.")
		c.Abort()
	}

	jwtWrapper := utils.JwtWrapper{
		SecretKey: "mySecretKey",
		Issuer:    "VolunteeryAuthService",
	}

	claims, err := jwtWrapper.ValidateToken(strings.TrimSpace(extractedToken[1]))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		c.Abort()
	}

	c.Set("userID", claims.UserID)
	c.Next()
}
