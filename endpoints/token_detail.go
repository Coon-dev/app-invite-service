package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	collection := configs.MongoClient.Database("pulseid").Collection("token")

	resp := models.TokenDetailResponse{
		Status: "not_found",
	}

	var token models.TokenList
	filter := bson.M{"token": req.Token}
	err := collection.FindOne(context.Background(), filter).Decode(&token)
	if err == mongo.ErrNoDocuments {
		configs.Clog.Printf("token not found: %+v", err)
		c.JSON(http.StatusOK, resp)
		return
	} else if err != nil {
		configs.Clog.Printf("Select mongo error: %+v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp.Status = token.Status
	if time.Now().After(token.ExpiredAt) {
		resp.Status = "inactive"
	}

	configs.Clog.Printf("[%+v] response: %+v", req.Token, resp)
	c.JSON(http.StatusOK, resp)
}
