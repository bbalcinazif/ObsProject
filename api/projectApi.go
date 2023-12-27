package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func getProjects(c *gin.Context) {
	var projects []Models.Project
	if err := Models.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Projeler görüntülenemedi",
		})
	}
	c.JSON(http.StatusOK, projects)

}

func deleteProjectByID(c *gin.Context) {
	id := c.Param("id")
	result := Models.DB.Delete(&Models.Project{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed",
		})
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Proje bulunamadı",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Proje başarıyla silindi.",
	})
}
func updateProject(c *gin.Context) {
	var updatedproject Models.Project
	id := c.Param("id")

	if err := c.Bind(&updatedproject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Geçersiz istek verisi",
		})
		return
	}

	project := Models.Project{}
	if err := Models.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Proje bulunamadı",
		})
		return
	}
	project.Name = updatedproject.Name
	project.ID = updatedproject.ID

	if err := Models.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Proje güncelleme başarısız",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Proje başarıyla güncellendi",
	})
}

// TODO lesson apideki gibi proje kaydında bir lessonproject kaydı yapacak ...
func signProject(c *gin.Context) {
	var data map[string]interface{}
	var project Models.Project
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

	project.Name = data["project_name"].(string)

	result := Models.DB.Create(&project)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create lesson",
		})
		return
	}

	var lessonproject Models.LessonProject
	lessonproject.ProjectID = project.ID
	lessonproject.LessonsID = uint(data["lessons_id"].(float64))

	resultls := Models.DB.Create(&lessonproject)
	if resultls.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Lessons project oluşturulamadı",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Başarılı",
	})

}

func ProjectApi(r *gin.RouterGroup) {
	r.GET("/getprojects", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getProjects)
	r.POST("/signproject", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, signProject)
	r.DELETE("/deleteprojectbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, deleteProjectByID)
	r.PUT("/updateprojectbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, updateProject)

}
