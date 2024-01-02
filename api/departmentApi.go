package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDepartments(c *gin.Context) {
	var departments []Models.Department
	if err := Models.DB.Find(&departments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bölümler Görüntülenemedi1",
		})
	}
	c.JSON(http.StatusOK, departments)
}
func signDepartment(c *gin.Context) {
	var Department Models.Department
	err := c.Bind(&Department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	result := Models.DB.Create(&Department)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Department",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Department oluşturuldu",
	})

}

func DepartmentApi(r *gin.RouterGroup) {
	r.GET("/getdepartments", MiddleWare.IsJwtValid, MiddleWare.IsManager, getDepartments)
	r.POST("/signdepartment", signDepartment)
}
