package Models

import "gorm.io/gorm"

type Notes struct {
	gorm.Model
	Vize1     int     `json:"vize_1"`
	Vize2     int     `json:"vize_2"`
	Final     int     `json:"final"`
	ProjectID uint    `json:"project_id"`
	Project   Project `json:"project"`
	ProjeNot  int     `json:"proje_not"`
	LessonsID uint    `json:"lessons_id"`
	Lessons   Lessons `json:"lessons"`
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	ISPass    *bool   `json:"is_pass"`
	//TODO harf notu kaydet
}
