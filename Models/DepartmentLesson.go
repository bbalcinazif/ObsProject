package Models

import "gorm.io/gorm"

type DepartmentLesson struct {
	gorm.Model
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
	LessonsID    uint       `json:"lessons_id"`
	Lessons      Lessons    `json:"lessons"`
}
