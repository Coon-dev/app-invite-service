package test

import (
	"net/http"
	"server/app-invite-service/models"
	"server/app-invite-service/services"
	"server/app-invite-service/utils"
	"testing"
)

func TestTokenDetailService(t *testing.T) {

	request := models.TokenDetailRequest{
		Token: "testtoken",
	}
	mockData := &MockDatabase{}

	//case 1: no record
	mockData.Case = "case1111"
	httpStatus1, resp1 := services.TokenDetailService(request, mockData)
	if httpStatus1 != http.StatusOK || resp1.Status != utils.StatusNotFound {
		t.Errorf("case 1 no record: expected http status 200 and token status: %s, found: http status: %d, token status: %s", utils.StatusNotFound, httpStatus1, resp1.Status)
	}

	//case 2: other error
	mockData.Case = "case2222"
	httpStatus2, resp2 := services.TokenDetailService(request, mockData)
	if httpStatus2 != http.StatusInternalServerError || resp2 != nil {
		t.Errorf("case 2 other: expected http status 500 and resp: nil, found: http status: %d, token: %s", httpStatus2, resp2)
	}

	//case 3: token expire
	mockData.Case = "case3333"
	httpStatus3, resp3 := services.TokenDetailService(request, mockData)
	if httpStatus3 != http.StatusOK || resp3.Status != utils.StatusInactive {
		t.Errorf("case 3 token expire: expected http status 200 and resp: %s, found: http status: %d, token: %s", utils.StatusInactive, httpStatus3, resp3)
	}

	//case 4: token active
	mockData.Case = "tokenActive"
	httpStatus4, resp4 := services.TokenDetailService(request, mockData)
	if httpStatus4 != http.StatusOK || resp4.Status != utils.StatusActive {
		t.Errorf("case 4 token active: expected http status 200 and resp: %s, found: http status: %d, token: %s", utils.StatusActive, httpStatus4, resp4.Status)
	}

}

func TestTokenDisableService(t *testing.T) {
	request := models.TokenDisableRequest{
		Token: "testtoken",
	}
	mockData := &MockDatabase{}

	//case 1: success
	mockData.Case = "case success"
	httpStatus1 := services.TokenDisableService(request, mockData)
	if httpStatus1 != http.StatusOK {
		t.Errorf("case 1 no record: expected http status 200, found: http status: %d", httpStatus1)
	}

	//case 2: database error
	mockData.Case = "case database error"
	httpStatus2 := services.TokenDisableService(request, mockData)
	if httpStatus2 != http.StatusInternalServerError {
		t.Errorf("case 2: database error: expected http status 500, found: http status: %d", httpStatus2)
	}

	//case 3: record not found
	mockData.Case = "case record not found"
	httpStatus3 := services.TokenDisableService(request, mockData)
	if httpStatus3 != http.StatusNotModified {
		t.Errorf("case 3: record not found: expected http status 304, found: http status: %d", httpStatus3)
	}
}

func TestTokenGenerateService(t *testing.T) {

	mockData := &MockDatabase{}

	//case 1: insert error
	mockData.Case = "insert1111"
	httpStatus1, resp1 := services.TokenGenerateService(mockData)
	if httpStatus1 != http.StatusInternalServerError || resp1 != nil {
		t.Errorf("case 1 no record: expected http status 500 and token: nil, found: http status: %d, token: %s", httpStatus1, resp1)
	}

	//case 2: success
	mockData.Case = "insert success"
	httpStatus2, resp2 := services.TokenGenerateService(mockData)
	if httpStatus2 != http.StatusOK {
		t.Errorf("case 2 success: expected http status 200, found: http status: %d", httpStatus2)
	}

	if resp2.Status != utils.StatusActive {
		t.Errorf("case 2 success: expected status='active', found: status: %s", resp2.Status)
	}

	if len(resp2.Token) < 6 || len(resp2.Token) > 12 {
		t.Errorf("case 2 success: expected token length between 6 and 12, found: length: %d", len(resp2.Token))
	}

	alphanumeric := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	token_rune := []rune(resp2.Token)
	for i := range token_rune {
		found := false
		for _, al := range alphanumeric {
			if token_rune[i] == al {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("case 2 success: expected all alphanumberic, found: %s, index %d", resp2.Token, i)
		}
	}
}

func TestTokenListService(t *testing.T) {

	mockData := &MockDatabase{}

	//case 1: database error
	mockData.Case = "database error"
	httpStatus1, resp1 := services.TokenListService(mockData)
	if httpStatus1 != http.StatusInternalServerError || resp1 != nil {
		t.Errorf("case 1 database error: expected http status 500 and token: nil, found: http status: %d, token: %s", httpStatus1, resp1)
	}

	//case 2: success
	mockData.Case = "success"
	httpStatus2, resp2 := services.TokenListService(mockData)
	if httpStatus2 != http.StatusOK {
		t.Errorf("case 2 success: expected http status 500, found: http status: %d", httpStatus2)
	}

	if resp2.TokenList[0].Status != utils.StatusActive {
		t.Errorf("case 2 success: expected token status 'active', found: token status: %s, token: %s", resp2.TokenList[0].Status, resp2.TokenList[0].Token)
	}

	if resp2.TokenList[1].Status != utils.StatusInactive {
		t.Errorf("case 2 success: expected token status 'inactive', found: token status: %s, token: %s", resp2.TokenList[1].Status, resp2.TokenList[1].Token)
	}

	if resp2.TokenList[2].Status != utils.StatusInactive {
		t.Errorf("case 2 success: expected token status 'inactive', found: token status: %s, token: %s", resp2.TokenList[2].Status, resp2.TokenList[2].Token)
	}
}
