package api

import (
	"ObsProject/Controllers"
	"ObsProject/MiddleWare"
	"ObsProject/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func userSignup(c *gin.Context) {
	var user Models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fields cannot be empty",
		})
		return
	}
	if user.Password != "" && user.Username != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash Password",
			})
		}
		user.Password = string(hash)
		result := Models.DB.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create User",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User created"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields cannot be empty",
		})
	}

}
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	result := Models.DB.Delete(&Models.User{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed",
		})
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User bulunamadı",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Kullanıcı başarıyla silindi",
	})
}

func getUsers(c *gin.Context) {
	var users []Models.User
	if err := Models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Kullanıcılar görüntülenemedi."})
	}
	c.JSON(http.StatusOK, users)
}
func UserApi(r *gin.RouterGroup) {
	r.POST("/loginuser", Controllers.LoginUser)
	r.DELETE("/deleteuserbyid/:id", MiddleWare.IsJwtValid, MiddleWare.IsManager, deleteUser)
	r.POST("/usersignup", MiddleWare.IsJwtValid, userSignup)
	r.GET("getusers", MiddleWare.IsJwtValid, MiddleWare.IsManager, getUsers)

}
