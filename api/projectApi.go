package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProjects(c *gin.Context) {
	var projects []Models.Project
	if err := Models.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Projeler görüntülenemedi",
		})
	}
	c.JSON(http.StatusOK, projects)

}

func DeleteProjectByID(c *gin.Context) {
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

func SignProject(c *gin.Context) {
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
	r.GET("/getprojects", MiddleWare.IsJwtValid, GetProjects)
	r.POST("/signproject", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, SignProject)
	r.DELETE("/deleteprojectbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsTeacher, DeleteProjectByID)
}
