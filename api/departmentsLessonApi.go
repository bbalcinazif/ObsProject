package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLessonsFromDepartmentID(c *gin.Context) {
	dID := c.Param("id")
	var dLessons []Models.DepartmentLesson

	if err := Models.DB.Where("department_id=?", dID).Find(&dLessons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Bölüm dersleri alınamadı."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"department_lessons": dLessons,
	})
}
func DepartmentLessonApi(r *gin.RouterGroup) {
	r.GET("/lessonsofdepartmentsid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, GetLessonsFromDepartmentID)
}
