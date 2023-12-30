package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID uint   `json:"post_id"`
	UserID uint   `json:"user_id"`
	Text   string `json:"text"`
	Likes  uint   `json:"likes"`
}
