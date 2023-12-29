package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	ImageURL string `json:"imageurl"`
}
