package services

import (
	"context"
	"math/rand"
	"net/http"
	"server/app-invite-service/configs"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TokenDetailService(request models.TokenDetailRequest) (int, *models.TokenDetailResponse) {
	collection := configs.MongoClient.Database("pulseid").Collection("token")

	resp := &models.TokenDetailResponse{
		Status: utils.StatusNotFound,
	}

	var token models.TokenList
	filter := bson.M{"token": request.Token}
	err := collection.FindOne(context.Background(), filter).Decode(&token)
	if err == mongo.ErrNoDocuments {
		configs.Clog.Printf("[%s] token not found: %+v", request.Token, err)
		return http.StatusOK, resp
	} else if err != nil {
		configs.Clog.Printf("[%s] Select mongo error: %+v", request.Token, err)
		return http.StatusInternalServerError, nil
	}

	resp.Status = token.Status
	if time.Now().After(token.ExpiredAt) {
		resp.Status = utils.StatusInactive
	}

	configs.Clog.Printf("[%s] response: %+v", request.Token, resp)
	return http.StatusOK, resp
}

func TokenDisableService(request models.TokenDisableRequest) int {
	collection := configs.MongoClient.Database("pulseid").Collection("token")
	filter := bson.M{"token": request.Token}
	updator := bson.M{"$set": bson.M{"status": utils.StatusInactive}}
	result, err := collection.UpdateOne(context.Background(), filter, updator)
	if err != nil {
		configs.Clog.Printf("[%s] Insert database error: %+v", request.Token, err)
		return http.StatusInternalServerError
	}
	if result.ModifiedCount <= 0 {
		configs.Clog.Printf("[%s] token not found", request.Token)
		return http.StatusNotModified
	}
	configs.Clog.Printf("[%s] update completed", request.Token)
	return http.StatusOK
}

func TokenGenerateService() (int, *models.TokenGenerateResponse) {
	newToken := utils.RandomToken(6 + rand.Intn(7))

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
		return http.StatusInternalServerError, nil
	}

	resp := &models.TokenGenerateResponse{
		Token:     newToken,
		Status:    utils.StatusActive,
		CreatedAt: tn,
		ExpiredAt: tn.AddDate(0, 0, 7),
	}
	configs.Clog.Printf("response: %+v", resp)
	return http.StatusOK, resp
}

func TokenListService() (int, *models.TokenListResponse) {
	//Select database
	var token []models.TokenList
	collection := configs.MongoClient.Database("pulseid").Collection("token")
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		configs.Clog.Printf("Select mongo error: %+v", err)
		return http.StatusInternalServerError, nil
	}

	err = cur.All(context.Background(), &token)
	if err != nil {
		configs.Clog.Printf("cursor to token: %+v", err)
		return http.StatusInternalServerError, nil
	}

	for a := range token {
		if time.Now().After(token[a].ExpiredAt) {
			token[a].Status = utils.StatusInactive
		}
	}
	if token == nil {
		token = make([]models.TokenList, 0)
	}
	resp := &models.TokenListResponse{
		TokenList: token,
	}
	configs.Clog.Printf("response: %+v", resp)
	return http.StatusOK, resp
}
