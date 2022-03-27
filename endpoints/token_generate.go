package endpoints

import (
	"math/rand"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"

	"github.com/gin-gonic/gin"
)

func TokenGenerateEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if auth != configs.AuthKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	newToken := randomToken(6 + rand.Intn(7))

	//Insert database

	resp := models.TokenGenerateResponse{
		Token: newToken,
	}
	configs.Clog.Printf("response: %+v", resp)
	c.JSON(http.StatusOK, resp)
}

func randomToken(length int) string {
	alphanumeric := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(b)
}
