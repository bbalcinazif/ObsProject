package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLessons(c *gin.Context) {
	var lessons []Models.Lesson
	if err := Models.DB.Find(&lessons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dersler Görüntülenemedi.",
		})
	}
	c.JSON(http.StatusOK, lessons)
}
func GetLessonsByUserID(c *gin.Context) {

	userID := c.Param("id")
	var userLessons []Models.Lesson

	// lessons
	if err := Models.DB.Where("user_id = ?", userID).Find(&userLessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ders bilgileri alınamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_lessons": userLessons,
	})
}

func LessonApi(r *gin.RouterGroup) {
	r.GET("/getlessons", MiddleWare.IsJwtValid, GetLessons)
	r.GET("departmentlessonsbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsStudent, GetLessonsByUserID)
}
