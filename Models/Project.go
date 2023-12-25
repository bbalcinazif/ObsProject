package Models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name"`
}
