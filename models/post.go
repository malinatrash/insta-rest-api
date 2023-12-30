package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	Caption    string    `json:"caption"`
	ImageURL   string    `json:"image_url"`
	VideoURL   string    `json:"video_url"`
	LikesCount uint      `json:"likes_count"`
	Comments   []Comment `json:"comments"`
}
