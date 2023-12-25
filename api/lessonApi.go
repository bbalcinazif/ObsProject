package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func getLessons(c *gin.Context) {
	var lessons []Models.Lessons
	if err := Models.DB.Find(&lessons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dersler Görüntülenemedi.",
		})
	}
	c.JSON(http.StatusOK, lessons)
}
func getLessonsByUserID(c *gin.Context) {

	userID := c.Param("id")
	var userLessons []Models.Lessons

	// lessons
	if err := Models.DB.Where("user_id = ?", userID).Find(&userLessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ders bilgileri alınamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_lessons": userLessons,
	})
}

func signLesson(c *gin.Context) {
	var data map[string]interface{}
	var lesson Models.Lessons
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to request body",
		})
		return
	}
	err = json.Unmarshal(byteData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unmarshal hatası",
		})
		return
	}
	//TODO user alacağız
	//lesson.DepartmentID = uint(data["department_id"].(float64))
	lesson.Name = data["lessons_name"].(string)
	lesson.UserID = uint(data["user_id"].(float64))
	// Lessons nesnesini kaydet
	result := Models.DB.Create(&lesson)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create lesson",
		})
		return
	}
	Models.DB.Where("name=?", data["lessons_name"]).First(&lesson)

	var departmentLesson Models.DepartmentLesson
	departmentLesson.LessonsID = lesson.ID
	departmentLesson.DepartmentID = uint(data["department_id"].(float64))
	fmt.Println("departmenlss:", departmentLesson)
	resultds := Models.DB.Create(&departmentLesson)
	if resultds.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Department lesson oluşturulamadı.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Başarılı",
	})
}

func LessonApi(r *gin.RouterGroup) {
	r.GET("/getlessons", MiddleWare.IsJwtValid /*Öğretmendersidsine göre*/, getLessons)
	r.GET("/lessonsbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsStudent, getLessonsByUserID)
	r.POST("/signlesson", MiddleWare.IsJwtValid, MiddleWare.IsManager, signLesson)
}
