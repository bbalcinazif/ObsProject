package Controllers

import (
	"ObsProject/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"time"
)

var jwtKey = []byte("gizliAnahtar")

func generateJWT(username string, isTeacher bool, isManager bool, ID uint) (string, error) {
	// Token generate
	token := jwt.New(jwt.SigningMethodHS256)

	// Token features
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 525600).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = username
	claims["is_teacher"] = isTeacher
	claims["is_manager"] = isManager
	claims["user_id"] = ID

	// Token signed
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func LoginUser(c *gin.Context) {
	var data map[string]interface{}
	byteData, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(byteData, &data)

	var user Models.User
	Models.DB.Where("username=?", data["username"]).First(&user)
	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"].(string)))
	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User login failed",
		})
	} else {
		token, err := generateJWT(user.Username, user.IsTeacher, user.IsManager, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
			return
		}
		c.SetCookie("token", token, 360000, "/", "", true, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Login Successful",
			"token":   token,
		})
	}

}
