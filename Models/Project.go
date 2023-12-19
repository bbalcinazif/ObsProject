package Models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name      string  `json:"name"`
	LessonsID uint    `json:"lessons_id"`
	Lessons   Lessons `json:"lessons"`
}
