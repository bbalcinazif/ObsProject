package Models

import "gorm.io/gorm"

type Lessons struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user"`
}
