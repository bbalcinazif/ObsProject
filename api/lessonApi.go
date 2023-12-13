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

func LessonApi(r *gin.RouterGroup) {
	r.GET("/getlessons", MiddleWare.IsJwtValid, GetLessons)
}
