package models

//easyjson:json
type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//easyjson:json
type ContactList []Contact