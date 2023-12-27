package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getNotesS(c *gin.Context) {
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
		fmt.Println("dersler listesi alınamadı")
	}
	fmt.Println("Department Lesson List:", departmentLessonList)

	var Notes []Models.Notes
	var lesson Models.Lessons
	var note Models.Notes

	for l := range departmentLessonList {
		Models.DB.Preload("User").Where("id = ?", departmentLessonList[l].LessonsID).First(&lesson)

		fmt.Println("lesson.UserID:", lesson.UserID)
		fmt.Println("userID:", userID)
		if lesson.UserID == userID.(uint) {
			fmt.Println("eeeeeeeeeeeeeeeeeeeeeee")
			if e := Models.DB.Preload("User").Preload("User.Department").Preload("Lessons").Preload("Lessons.User").Preload("Lessons.User.Department").Where("lessons_id = ?", lesson.ID).Where("user_id", userID).First(&Notes).Error; e != nil {
				Notes = append(Notes, note)
			} else {
				c.JSON(http.StatusOK, Notes)
			}
		}
	}

}
func getNotesT(c *gin.Context) {
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
		fmt.Println("dersler listesi alınamadı")
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
func CheckPass(c *gin.Context) {
	userID := c.Param("id")
	var Notes Models.Notes
	var dersBasarili int

	if err := Models.DB.Where("user_id=?", userID).Find(&Notes).Error; err != nil { //buraya bir where daha gelecek ve departmentid kontrol edecek.
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not kayıtları bulunamadı",
		})
	} else {
		vize1 := Notes.Vize1
		vize2 := Notes.Vize2
		final := Notes.Final
		projectnot := Notes.ProjeNot

		dersBasarili = int(float64(vize1)*0.15 + float64(vize2)*0.15 + float64(final)*0.40 + float64(projectnot)*0.3)
		if dersBasarili >= 50 && dersBasarili < 60 {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlandı...!!  Ders Harf Notu:DC"})
		} else if dersBasarili >= 60 && dersBasarili < 70 {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlandı...!! Ders Harf Notu:CC"})
		} else if dersBasarili >= 70 && dersBasarili < 80 {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlandı...!! Ders Harf Notu:BB"})
		} else if dersBasarili >= 80 && dersBasarili < 90 {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlandı...!! Ders Harf Notu:BA"})
		} else if dersBasarili >= 90 && dersBasarili < 100 {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlandı...!! Ders Harf Notu:AA"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Ders başarıyla tamamlanamadı...!! Ders Harf Notu:FF "})
		}
	}
	var newIsPass bool
	if dersBasarili >= 50 {
		newIsPass = true
	} else {
		newIsPass = false
	}
	if err := Models.DB.Model(&Notes).Update("is_pass", newIsPass).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ispass değeri güncellenirken hata oluştu",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ispass değeri başarıyla güncellendi",
	})

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
func deleteNotByID(c *gin.Context) {
	id := c.Param("id")
	result := Models.DB.Delete(&Models.Notes{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed",
		})
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not bulunamadı",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"message": "Not başarıyla silindi",
	})

}

func updateNot(c *gin.Context) {
	var updatednote Models.Notes
	id := c.Param("id")
	if err := c.Bind(&updatednote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Geçersiz istek",
		})
		return
	}
	note := Models.Notes{}
	if err := Models.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not bulunamadı",
		})
		return
	}
	note.Vize1 = updatednote.Vize1
	note.Vize2 = updatednote.Vize2
	note.Final = updatednote.Final

	if err := Models.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note güncelleme başarısız",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Not güncellendi",
	})

}

func signNote(c *gin.Context) {
	//öğretmen lesson idsiyle check yapacak...
	var note Models.Notes
	err := c.Bind(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	if note.Vize1 >= 0 && note.Vize1 <= 100 && note.Vize2 >= 0 && note.Vize2 <= 100 && note.Final >= 0 && note.Final <= 100 && note.ProjeNot >= 0 && note.ProjeNot <= 100 {
		result := Models.DB.Create(&note)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create note",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Not girişi başarılı",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Hatalı not girişi yaptınız ! Not değerleri 0 - 100 arası olmalıdır.",
		})
	}

}

func NoteApi(r *gin.RouterGroup) {
	r.GET("/getnotesbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getNoteByID)
	r.GET("/checkpass/:id", MiddleWare.IsJwtValid, CheckPass)
	r.POST("signnote", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, signNote)
	r.GET("/getnotest", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getNotesT)
	r.GET("/getnotes", MiddleWare.IsJwtValid, MiddleWare.IsStudent, getNotesS)
	r.DELETE("/deletenotebyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, deleteNotByID)
	r.PUT("/updatenotebyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, updateNot)
}
