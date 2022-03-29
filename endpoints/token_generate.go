package endpoints

import (
	"net/http"
	"server/app-invite-service/services"

	"github.com/gin-gonic/gin"
)

func TokenGenerateEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if !services.AuthService(auth) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(services.TokenGenerateService())
}
