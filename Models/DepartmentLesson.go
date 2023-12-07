package Models

import "gorm.io/gorm"

type DepartmentLesson struct {
	gorm.Model
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
	LessonID     uint       `json:"lesson_id"`
	Lesson       Lesson     `json:"lesson"`
}
