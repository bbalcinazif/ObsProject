package MiddleWare

import (
	"ObsProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func IsJwtValid(c *gin.Context) {
	tokenString, _ := c.Request.Cookie("token")

	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("gizliAnahtar"), nil
	})

	if err != nil {
		fmt.Println("JWT Decode failed ", err)
		c.AbortWithStatus(401)
		return

	} else {
		//JWT son kullanım süresi

		if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expirationTime := time.Unix(int64(Claims["exp"].(float64)), 0)
			currentTime := time.Now()

			if currentTime.Before(expirationTime) {
				c.Next()

			} else {
				c.AbortWithStatus(401)
			}
		} else {
			c.AbortWithStatus(401)
		}
	}

}
func IsTeacher(c *gin.Context) {
	// Kullanıcı bilgisini çekme
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(401)
		return
	}

	// user'dan IsTeacher bilgisini al
	userModel, ok := user.(*Models.User)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	// İsteği yapan kullanıcı öğretmen mi kontrol etme
	if userModel.IsTeacher != nil && *userModel.IsTeacher == true {
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Only teacher acces to this request",
		})
	}
}
func IsStudent() {

}
