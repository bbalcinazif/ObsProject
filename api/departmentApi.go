package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDepartments(c *gin.Context) {
	var departments []Models.Department
	if err := Models.DB.Find(&departments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bölümler Görüntülenemedi1",
		})
	}
	c.JSON(http.StatusOK, departments)
}
func DepartmentApi(r *gin.RouterGroup) {
	r.GET("/getdepartments", MiddleWare.IsJwtValid, GetDepartments)
}
