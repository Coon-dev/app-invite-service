package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"

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

	//Select database

	statusMock := [2]string{"active", "inactive"}
	resp := models.TokenDetailResponse{
		Status: statusMock[rand.Intn(2)],
	}
	configs.Clog.Printf("[%+v] response: %+v", req.Token, resp)
	c.JSON(http.StatusOK, resp)
}
