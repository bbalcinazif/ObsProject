package api

import (
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLesson(c *gin.Context) {
	var lessons []Models.Lesson
	if err := Models.DB.Find(&lessons).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dersler Görüntülenemedi",
		})

	}
	c.JSON(http.StatusOK, lessons)
}
func LessonApi(r *gin.RouterGroup) {
	r.GET("/getlessons", GetLesson)

}
