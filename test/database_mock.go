package test

import (
	"errors"
	"server/app-invite-service/models"
	"server/app-invite-service/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MockDatabase struct {
	Case string
}

func (mg *MockDatabase) FindOne(filter interface{}, result interface{}) error {

	r, ok := result.(*models.TokenList)
	if !ok {
		panic("mock response error")
	}

	switch mg.Case {
	case "case1111":
		return mongo.ErrNoDocuments
	case "case2222":
		return errors.New("other error")
	case "case3333":
		*r = models.TokenList{
			Token:     "asdf12312",
			Status:    "active",
			CreatedAt: time.Now().AddDate(0, 0, -8),
			ExpiredAt: time.Now().AddDate(0, 0, -1),
		}
		return nil
	}

	*r = models.TokenList{
		Token:     "asdf12312",
		Status:    "active",
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().AddDate(0, 0, 7),
	}
	return nil
}

func (mg *MockDatabase) FindAll(filter interface{}, result interface{}) error {
	switch mg.Case {
	case "database error":
		return errors.New("database error")
	}
	r, ok := result.(*[]models.TokenList)
	if !ok {
		panic("mock response error")
	}

	tokenList := make([]models.TokenList, 3)
	tokenList[0] = models.TokenList{
		Token:     "tokenOne1",
		Status:    utils.StatusActive,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().AddDate(0, 0, 7),
	}

	tokenList[1] = models.TokenList{
		Token:     "tokenTwo2",
		Status:    utils.StatusInactive,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().AddDate(0, 0, 7),
	}

	tokenList[2] = models.TokenList{
		Token:     "tokenThree3",
		Status:    utils.StatusActive,
		CreatedAt: time.Now().AddDate(0, 0, -8),
		ExpiredAt: time.Now().AddDate(0, 0, -1),
	}

	*r = tokenList

	return nil
}

func (mg *MockDatabase) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	switch mg.Case {
	case "case database error":
		return nil, errors.New("error")
	case "case record not found":
		return &mongo.UpdateResult{ModifiedCount: 0}, nil
	}

	return &mongo.UpdateResult{ModifiedCount: 1}, nil
}

func (mg *MockDatabase) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {

	switch mg.Case {
	case "insert1111":
		return nil, errors.New("error")
	}

	return &mongo.InsertOneResult{}, nil
}
