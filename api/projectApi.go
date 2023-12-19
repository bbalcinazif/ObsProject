package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
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

func signProject(c *gin.Context) {
	var project Models.Project
	err := c.Bind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "a cannot be empty",
		})
		return
	}

	if project.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Proje adı boş bırakılamaz",
		})
		return
	}

	result := Models.DB.Create(&project)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project created",
	})
}

func ProjectApi(r *gin.RouterGroup) {
	r.GET("/getprojects", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, getProjects)
	r.POST("/signproject", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, signProject)
	r.DELETE("/deleteprojectbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, deleteProjectByID)
	r.PUT("/updateprojectbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, updateProject)

}
