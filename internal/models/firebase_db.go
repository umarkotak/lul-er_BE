package models

type FirebaseDB struct {
	Users map[string]User `json:"users"`
}
