package validators

import (
	"net/http"

	"github.com/bentanjunrong/Volunteer-Board-CCSGP-Backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RegisterValidator(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
	}
	c.Next()
}
