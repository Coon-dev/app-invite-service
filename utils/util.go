package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/app-invite-service/configs"
)

func Request(method string, url string, header map[string]string, reqObj interface{}, resObj interface{}) error {
	// Marshal request object in to bytes
	marshal, err := json.Marshal(reqObj)
	if err != nil {
		return err
	}

	// Create HTTP request object from bytes
	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}

	if header != nil {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}

	// Post JSON request object
	response, err := configs.HTTPCilent.Do(request)
	if err != nil {
		return err
	}

	// Ensure that there is no error
	if response.StatusCode != http.StatusOK {
		configs.Clog.Printf("response status: %+v", response.StatusCode)
		err := fmt.Errorf("Error: HTTP-CODE=%d", response.StatusCode)
		return err
	}

	// Read response text message
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Parse response text message as specified buffer
	if resObj != nil {
		if err := json.Unmarshal([]byte(body), &resObj); err != nil {
			return err
		}
	}
	configs.Clog.Printf("response body: %+v", resObj)

	return nil
}
