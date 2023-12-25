package Models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	Name         string     `json:"name"`
	Surname      string     `json:"surname"`
	Mail         string     `json:"mail"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
	ProjectID    *uint      `json:"project_id"`
	Project      Project    `json:"project"`
	IsTeacher    bool       `json:"is_teacher"`
	IsManager    bool       `json:"is_manager"`
}
