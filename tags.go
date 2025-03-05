package yonoma

import "fmt"

// TagParams for creating/updating tags
type TagParams struct {
	Name string `json:"name,omitempty"`
}

// Tag represents a Yonoma tag
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List all tags
func ListTags() ([]Tag, error) {
	var tags []Tag
	err := request("GET", "tags/list", nil, &tags)
	return tags, err
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
