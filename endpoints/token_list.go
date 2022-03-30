package endpoints

import (
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/services"

	"github.com/gin-gonic/gin"
)

func TokenListEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if !services.AuthService(auth) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	mgdb := &services.MongoDatabase{
		Collection: configs.MongoClient.Database("pulseid").Collection("token"),
	}
	c.JSON(services.TokenListService(mgdb))
}
