package endpoints

import (
	"context"
	"math/rand"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"
	"time"

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

	tn := time.Now()
	insert := models.TokenList{
		Token:     newToken,
		Status:    utils.StatusActive,
		CreatedAt: tn,
		ExpiredAt: tn.AddDate(0, 0, 7),
	}

	//Insert database
	collection := configs.MongoClient.Database("pulseid").Collection("token")
	_, err := collection.InsertOne(context.Background(), insert)
	if err != nil {
		configs.Clog.Printf("Insert database error: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := models.TokenGenerateResponse{
		Token:     newToken,
		Status:    utils.StatusActive,
		CreatedAt: tn,
		ExpiredAt: tn.AddDate(0, 0, 7),
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
