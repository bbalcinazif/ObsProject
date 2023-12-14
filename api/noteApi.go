package api

import (
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotes(c *gin.Context) {

	//öğretmen departmanının notlarını görebilecek
}

func GetNoteByID(c *gin.Context) {
	noteID := c.Param("id")

	if noteID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID zorunludur",
		})
		return
	}
	var note Models.Notes

	if err := Models.DB.First(&note, noteID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not bulunamadı",
		})
		return
	}

	c.JSON(http.StatusOK, note)
}
func SignNote(c *gin.Context) {
	var note Models.Notes
	err := c.Bind(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	result := Models.DB.Create(&note)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to join note",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Not girişi başarılı",
	})
}

func NoteApi(r *gin.RouterGroup) {
	r.GET("/getnotesbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsStudent /*DepartmanIDCheck    */, GetNoteByID)
	r.POST("signnote", MiddleWare.IsJwtValid, MiddleWare.IsTeacher /* DepartmanIDCheck  */, SignProject)

}
