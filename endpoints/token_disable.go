package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"

	"github.com/gin-gonic/gin"
)

func TokenDisableEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if auth != configs.AuthKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var req models.TokenDisableRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		configs.Clog.Printf("cannot decode body: %+v", c.Request.Body)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Update database

	statusMock := [2]int{http.StatusOK, http.StatusBadRequest}

	c.JSON(statusMock[rand.Intn(2)], nil)
}
