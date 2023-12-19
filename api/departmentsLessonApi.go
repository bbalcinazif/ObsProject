package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getLessonsFromDepartmentID(c *gin.Context) {
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
func signDepartmentLesson(c *gin.Context) {
	var departmentLesson Models.DepartmentLesson
	err := c.Bind(&departmentLesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	result := Models.DB.Create(&departmentLesson)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create DepartmentLesson",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Department Lesson oluşturuldu",
	})

}

func DepartmentLessonApi(r *gin.RouterGroup) {
	r.GET("/lessonsofdepartmentsid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getLessonsFromDepartmentID)
	r.POST("/signdlesson", MiddleWare.IsJwtValid, MiddleWare.IsManager, signDepartmentLesson)

}
