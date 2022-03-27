package endpoints

import (
	"net/http"
	"server/app-invite-service/configs"

	"github.com/gin-gonic/gin"
)

func PingEndpoint(c *gin.Context) {
	configs.Clog.Println("[ping] pong ja")
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
