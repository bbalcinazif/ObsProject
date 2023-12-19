package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getNotes(c *gin.Context) {
	tokenString, _ := c.Request.Cookie("token")
	userID := MiddleWare.GetUserInToken(tokenString.Value)
	fmt.Println("userID:", userID)
	var user Models.User
	if err := Models.DB.Preload("Department").Where("id = ?", userID).First(&user).Error; err != nil {
		fmt.Println("user alınamadı..")
	}

	fmt.Println("USER:", user)
	fmt.Println("department.id:", user.DepartmentID)

	var departmentLessonList []Models.DepartmentLesson
	if er := Models.DB.Preload("Lessons").Where("department_id = ?", user.DepartmentID).Find(&departmentLessonList).Error; er != nil {
		fmt.Println("dersler listesini alamadım")
	}
	fmt.Println("Department Lesson List:", departmentLessonList)

	var notes []Models.Notes
	var lesson Models.Lessons
	for l := range departmentLessonList {
		Models.DB.Preload("User").Where("id = ?", departmentLessonList[l].LessonsID).First(&lesson)

		fmt.Println("lesson.UserID:", lesson.UserID)
		fmt.Println("userID:", userID)
		if lesson.UserID == userID.(uint) {
			fmt.Println("eeeeeeeeeeeeeeeeeeeeeee")
			if e := Models.DB.Preload("User").Preload("User.Department").Preload("Lessons").Preload("Lessons.User").Preload("Lessons.User.Department").Where("lessons_id = ?", lesson.ID).Find(&notes).Error; e != nil {
				fmt.Println("notlar getirilemedi.")
			} else {
				c.JSON(http.StatusOK, notes)
			}
		}
	}

	fmt.Println("NOTES:", notes)
}

func getNoteByID(c *gin.Context) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID zorunludur",
		})
		return
	}
	var notes Models.Notes

	if err := Models.DB.Where("user_id=?", userID).Find(&notes).Error; err != nil { //buraya bir where daha gelecek ve departmentid kontrol edecek.
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not bulunamadı",
		})
		return
	}

	c.JSON(http.StatusOK, notes)
}
func signNote(c *gin.Context) {
	var note Models.Notes
	err := c.Bind(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	result := Models.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Not girişi başarılı",
	})
}

func NoteApi(r *gin.RouterGroup) {
	r.GET("/getnotesbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher /*DepartmanIDCheck    */, getNoteByID)
	r.POST("signnote", MiddleWare.IsJwtValid, MiddleWare.IsTeacher /* DepartmanIDCheck  */, signNote)
	r.GET("/getnotest", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getNotes)

}
