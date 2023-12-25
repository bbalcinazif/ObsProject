package Models

import "gorm.io/gorm"

type LessonProject struct {
	gorm.Model
	LessonsID uint    `json:"lessons_id"`
	Lessons   Lessons `json:"lessons"`
	ProjectID uint    `json:"project_id"`
	Project   Project `json:"project"`
}
