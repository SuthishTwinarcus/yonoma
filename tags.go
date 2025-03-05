package yonoma

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TagParams for creating/updating tags
type TagParams struct {
	Name string `json:"name,omitempty"`
}

// Tag represents a Yonoma tag
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Tags struct {
	client TagsYonomaClient
}
type TagsYonomaClient struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func (yc *TagsYonomaClient) Request(method, endpoint string, data interface{}) (map[string]interface{}, error) {
	url := yc.baseURL + endpoint
	var requestBody []byte
	var err error

	if data != nil {
		requestBody, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+yc.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := yc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

// List all tags
//
//	func ListTags() ([]Tag, error) {
//		//var tags []Tag
//		var tags []Tag = []Tag{}
//		err := request("GET", "tags/list", nil, &tags)
//		return tags, err
//	}
func (t *Tags) ListTags() (map[string]any, error) {
	fmt.Println("Created List")
	//var tags []Tag
	//var tags []Tag = []Tag{}
	//err := request("GET", "tags/list", nil, &tags)
	return t.client.Request("GET", "tags/list", nil)
}

// New creates a new tag
func NewTag(params *TagParams) (*Tag, error) {
	var tag Tag
	err := request("POST", "/tags", params, &tag)
	return &tag, err
}

// Retrieve a tag
func RetrieveTag(tagID string) (*Tag, error) {
	url := fmt.Sprintf("/tags/%s", tagID)
	var tag Tag
	err := request("GET", url, nil, &tag)
	return &tag, err
}

// Update a tag
func UpdateTag(tagID string, params *TagParams) error {
	url := fmt.Sprintf("/tags/%s", tagID)
	return request("PATCH", url, params, nil)
}

// Delete a tag
func DeleteTag(tagID string) error {
	url := fmt.Sprintf("/tags/%s", tagID)
	return request("DELETE", url, nil, nil)
}
