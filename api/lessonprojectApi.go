package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func getProject(c *gin.Context) {
	tokenString, _ := c.Request.Cookie("token")
	userID := MiddleWare.GetUserInToken(tokenString.Value)
	fmt.Println("userID:", userID)
	var user Models.User
	if err := Models.DB.Preload("Department").Where("id=?", userID).First(&user).Error; err != nil {
		fmt.Println(" User alınamadı")
	}
	fmt.Println("User:", user)
	fmt.Println("DepartmentID", user.DepartmentID)

	var LessonProjectList []Models.DepartmentLesson

	if er := Models.DB.Preload("Project").Where("lessons_id=?", user.ProjectID).Find(&LessonProjectList).Error; er != nil {
		fmt.Println("Proje listesi alınamadı")
	}
	fmt.Println("Lesson List:", LessonProjectList)

}

func LessonProjectApi(r *gin.RouterGroup) {
	r.GET("/getlessonprojects", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getProject)
}
