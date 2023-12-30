package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}
