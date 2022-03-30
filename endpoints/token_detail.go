package endpoints

import (
	"encoding/json"
	"log"
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
		log.Printf("cannot decode body: %+v\n", c.Request.Body)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	mgdb := &services.MongoDatabase{
		Collection: configs.MongoClient.Database("pulseid").Collection("token"),
	}
	c.JSON(services.TokenDetailService(req, mgdb))
}
