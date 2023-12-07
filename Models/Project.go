package Models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name     string `json:"name"`
	LessonID uint   `json:"lesson_id"`
	Lesson   Lesson `json:"lesson"`
}
