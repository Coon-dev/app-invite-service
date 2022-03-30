package services

import (
	"log"
	"math/rand"
	"net/http"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TokenDetailService(request models.TokenDetailRequest, database models.Database) (int, *models.TokenDetailResponse) {

	resp := &models.TokenDetailResponse{
		Status: utils.StatusNotFound,
	}

	var token models.TokenList
	filter := bson.M{"token": request.Token}
	err := database.FindOne(filter, &token)
	if err == mongo.ErrNoDocuments {
		log.Printf("[%s] token not found: %+v\n", request.Token, err)
		return http.StatusOK, resp
	} else if err != nil {
		log.Printf("[%s] Select mongo error: %+v\n", request.Token, err)
		return http.StatusInternalServerError, nil
	}

	resp.Status = token.Status
	if time.Now().After(token.ExpiredAt) {
		resp.Status = utils.StatusInactive
	}

	log.Printf("[%s] response: %+v\n", request.Token, resp)
	return http.StatusOK, resp
}

func TokenDisableService(request models.TokenDisableRequest, database models.Database) int {
	filter := bson.M{"token": request.Token}
	updator := bson.M{"$set": bson.M{"status": utils.StatusInactive}}
	result, err := database.UpdateOne(filter, updator)
	if err != nil {
		log.Printf("[%s] Insert database error: %+v\n", request.Token, err)
		return http.StatusInternalServerError
	}
	if result.ModifiedCount <= 0 {
		log.Printf("[%s] token not found\n", request.Token)
		return http.StatusNotModified
	}
	log.Printf("[%s] update completed\n", request.Token)
	return http.StatusOK
}

func TokenGenerateService(database models.Database) (int, *models.TokenGenerateResponse) {
	newToken := utils.RandomToken(6 + rand.Intn(7))

	tn := time.Now()
	insert := models.TokenList{
		Token:     newToken,
		Status:    utils.StatusActive,
		CreatedAt: tn,
		ExpiredAt: tn.AddDate(0, 0, 7),
	}

	_, err := database.InsertOne(insert)
	if err != nil {
		log.Printf("Insert database error: %+v\n", err)
		return http.StatusInternalServerError, nil
	}

	resp := &models.TokenGenerateResponse{
		Token:     newToken,
		Status:    utils.StatusActive,
		CreatedAt: tn,
		ExpiredAt: tn.AddDate(0, 0, 7),
	}
	log.Printf("response: %+v\n", resp)
	return http.StatusOK, resp
}

func TokenListService(database models.Database) (int, *models.TokenListResponse) {

	var token []models.TokenList
	err := database.FindAll(bson.M{}, &token)
	if err != nil {
		log.Printf("database FindAll error: %+v\n", err)
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
	log.Printf("response: %+v\n", resp)
	return http.StatusOK, resp
}
