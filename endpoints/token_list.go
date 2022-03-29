package endpoints

import (
	"context"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TokenListEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()

	auth := c.Request.Header.Get("Authorization")
	if auth != configs.AuthKey {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//Select database
	var token []models.TokenList
	collection := configs.MongoClient.Database("pulseid").Collection("token")
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		configs.Clog.Printf("Select mongo error: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = cur.All(context.Background(), &token)
	if err != nil {
		configs.Clog.Printf("Select mongo error: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for a := range token {
		if time.Now().After(token[a].ExpiredAt) {
			token[a].Status = utils.StatusInactive
		}
	}
	resp := models.TokenListResponse{
		TokenList: token,
	}
	configs.Clog.Printf("response: %+v", resp)
	c.JSON(http.StatusOK, resp)
}
