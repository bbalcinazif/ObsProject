package main

import (
	"ObsProject/Models"
	"ObsProject/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:Nazif.57@tcp(localhost:3306)/obsProject?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}
	Models.DB = db
	err = db.AutoMigrate(&Models.User{}, &Models.Project{}, &Models.Department{})
	if err != nil {
		fmt.Println(err.Error())
	}

	r := gin.Default()
	group := r.Group("/obsApi")

	api.UserApi(group)
	api.DepartmentApi(group)
	api.ProjectApi(group)
	api.LessonApi(group)

	r.Run(":8080")
}
