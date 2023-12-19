package MiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GetUserInToken(tokenString string) interface{} {
	//tokenString, err := c.Request.Cookie("token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("gizliAnahtar"), nil
	})

	if err != nil {
		fmt.Println("JWT Decode failed ", err)
		//return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		currentTime := time.Now()

		if currentTime.Before(expirationTime) {
			userID := int(claims["user_id"].(float64))
			return uint(userID)
		} else {
			return nil
		}
	}
	return nil
}

func IsTeacher(c *gin.Context) {

	tokenString, _ := c.Request.Cookie("token")
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("gizliAnahtar"), nil
	})

	if err != nil {
		fmt.Println("JWT Decode failed ", err)
		c.AbortWithStatus(401)
		return
	} else {
		if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			isTeacher := Claims["is_teacher"].(bool)

			if isTeacher {
				c.Next()

			} else {
				c.AbortWithStatus(401)
			}
		} else {
			c.AbortWithStatus(401)
		}
	}

}
func IsStudent(c *gin.Context) {
	tokenString, _ := c.Request.Cookie("token")
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("gizliAnahtar"), nil
	})
	if err != nil {
		fmt.Println("JWT Decode failed", err)
		c.AbortWithStatus(401)
		return
	} else {
		if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			isStudent := Claims["is_teacher"].(bool)

			if !isStudent {
				c.Next()
			} else {
				c.AbortWithStatus(401)
			}
		} else {
			c.AbortWithStatus(401)
		}
	}
}
func IsManager(c *gin.Context) {
	tokenString, _ := c.Request.Cookie("token")
	token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte("gizliAnahtar"), nil
	})
	if err != nil {
		fmt.Println("JWT Decode failed", err)
		c.AbortWithStatus(401)
		return
	} else {
		if Claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			isManager := Claims["is_manager"].(bool)

			if isManager {
				c.Next()
			} else {
				c.AbortWithStatus(401)
			}
		} else {
			c.AbortWithStatus(401)
		}
	}
}
