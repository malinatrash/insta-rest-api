package models

import "time"

type Session struct {
	BaseModel
	Session        string    `json:"session"`
	ExpireDateTime time.Time `json:"expire_date_time"`
	UserRefer      int
	User           User `gorm:"foreignKey:UserRefer"`
}

type SessionResponse struct {
	Session string `json:"session"`
}
