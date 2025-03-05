package yonoma

import "fmt"

// ContactParams for creating/updating contacts
type ContactParams struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// Contact represents a Yonoma contact
type Contact struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// New creates a new contact
func NewContact(params *ContactParams) (*Contact, error) {
	var contact Contact
	err := request("POST", "contacts/%s/create", params, &contact)
	return &contact, err
}

// Unsubscribe a contact
func UnsubscribeContact(contactID string) error {
	url := fmt.Sprintf("/contacts/%s/unsubscribe", contactID)
	return request("PATCH", url, nil, nil)
}

// Add a tag to a contact
func AddContactTag(contactID, tagID string) error {
	url := fmt.Sprintf("/contacts/%s/tags/%s", contactID, tagID)
	return request("POST", url, nil, nil)
}

// Remove a tag from a contact
func RemoveContactTag(contactID, tagID string) error {
	url := fmt.Sprintf("/contacts/%s/tags/%s", contactID, tagID)
	return request("DELETE", url, nil, nil)
}
