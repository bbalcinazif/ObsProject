package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotesT(c *gin.Context) {
	userID := MiddleWare.GetUserInToken(c)
	var user Models.User

	if err := Models.DB.Where("user_id=?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Kullanıcı bilgileri hatalı",
		})
		return
	}
	var dLessons []Models.DepartmentLesson
	if err := Models.DB.Where("department_id=?", user.DepartmentID).Find(&dLessons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Hatalı",
		})
	}
	var lesson Models.Lesson
	var notes []Models.Notes
	for i := range dLessons {
		Models.DB.Where("lesson_id=?", dLessons[i].LessonID).First(&lesson)
		if lesson.UserID == userID {
			Models.DB.Where("lesson_id=?", dLessons[i].LessonID).Find(&notes)
			c.JSON(http.StatusOK, notes)
		} else {
			if i == len(dLessons)-1 {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "A Failed",
				})
			}
		}

	}

}

//func GetNotesS(c *gin.Context) {
//	userID := MiddleWare.GetUserInToken(c)
//	var user Models.User
//	var dLessons Models.DepartmentLesson
//	if err := Models.DB.Where("user_id=?", userID).Where("department_id=?", dLessons).First(&user).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"message": "Kullanıcı bilgileri hatalı",
//		})
//		return
//	}
//
//}

func GetNoteByID(c *gin.Context) {
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
func SignNote(c *gin.Context) {
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
			"error": "Failed to join note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Not girişi başarılı",
	})
}

func NoteApi(r *gin.RouterGroup) {
	r.GET("/getnotesbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher /*DepartmanIDCheck    */, GetNoteByID)
	r.POST("signnote", MiddleWare.IsJwtValid, MiddleWare.IsTeacher /* DepartmanIDCheck  */, SignNote)
	r.GET("/getnotes", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, GetNotesT)

}
