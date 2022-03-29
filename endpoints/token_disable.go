package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	collection := configs.MongoClient.Database("pulseid").Collection("token")
	filter := bson.M{"token": req.Token}
	updator := bson.M{"$set": bson.M{"status": utils.StatusInactive}}
	result, err := collection.UpdateOne(context.Background(), filter, updator)
	if err != nil {
		configs.Clog.Printf("Insert database error: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if result.ModifiedCount <= 0 {
		c.JSON(http.StatusNotModified, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
