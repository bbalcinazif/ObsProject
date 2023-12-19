package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
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
	var lesson Models.Lessons
	err := c.Bind(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}

	result := Models.DB.Create(&lesson)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create lesson",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ders kayıt başarılı",
	})
}
func LessonApi(r *gin.RouterGroup) {
	r.GET("/getlessons", MiddleWare.IsJwtValid /*Öğretmendersidsine göre*/, getLessons)
	r.GET("/lessonsbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsStudent, getLessonsByUserID)
	r.POST("/signlesson", MiddleWare.IsJwtValid, MiddleWare.IsManager, signLesson)
}
