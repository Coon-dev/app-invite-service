package endpoints

import (
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"

	"github.com/gin-gonic/gin"
)

func TokenListEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if auth != configs.AuthKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//Select database

	resp := models.TokenListResponse{
		TokenList: make([]models.TokenList, 0),
	}
	configs.Clog.Printf("response: %+v", resp)
	c.JSON(http.StatusOK, resp)
}
