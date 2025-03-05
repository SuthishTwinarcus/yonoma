package yonoma

import "fmt"

// ListParams for creating/updating lists
type ListParams struct {
	Name string `json:"name,omitempty"`
}

// List represents a Yonoma list
type List struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List all lists
func ListLists() ([]List, error) {
	var lists []List
	err := request("GET", "lists/list", nil, &lists)
	return lists, err
}

// New creates a new list
func NewList(params *ListParams) (*List, error) {
	var list List
	err := request("POST", "lists/create", params, &list)
	return &list, err
}

// Retrieve a list
func RetrieveList(listID string) (*List, error) {
	url := fmt.Sprintf("lists/", listID)
	var list List
	err := request("GET", url, nil, &list)
	return &list, err
}

// Update a list
func UpdateList(listID string, params *ListParams) error {
	url := fmt.Sprintf("lists/" + listID + "/update")
	return request("POST", url, params, nil)
}

// Delete a list
func DeleteList(listID string) error {
	url := fmt.Sprintf("lists/" + listID + "/delete")
	return request("POST", url, nil, nil)
}
