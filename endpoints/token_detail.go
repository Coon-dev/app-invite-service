package endpoints

import (
	"encoding/json"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/services"

	"github.com/gin-gonic/gin"
)

func TokenDetailEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	var req models.TokenDetailRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		configs.Clog.Printf("cannot decode body: %+v", c.Request.Body)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(services.TokenDetailService(req))
}
