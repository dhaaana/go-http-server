package models

type User struct {
	ID       int    `json:"id"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
