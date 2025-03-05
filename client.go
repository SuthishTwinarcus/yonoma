package yonoma

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Key string // API Key

const baseURL = "http://localhost:8080/v1/"

func request(method, endpoint string, params interface{}, v interface{}) error {
	if Key == "" {
		return errors.New("API Key is required. Set yonoma.Key before making requests")
	}

	url := baseURL + endpoint

	var body []byte
	var err error
	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+Key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", responseBody)
	}

	return json.Unmarshal(responseBody, v)
}
