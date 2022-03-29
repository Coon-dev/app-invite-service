package endpoints

import (
	"encoding/json"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/services"

	"github.com/gin-gonic/gin"
)

func TokenDisableEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if !services.AuthService(auth) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var req models.TokenDisableRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		configs.Clog.Printf("cannot decode body: %+v", c.Request.Body)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(services.TokenDisableService(req))
}
