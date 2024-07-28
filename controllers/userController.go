package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/authentication/dto"
	"example.com/authentication/initializers"
	"example.com/authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get data from req body
	var userData dto.UserDTO
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Bind Data",
			"error":   err.Error(),
		})
		return
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Hash Password",
			"error":   err.Error(),
		})
		return
	}

	// Insert into Database
	user := models.User{
		Name:     userData.Name,
		Username: userData.Username,
		Password: string(hashedPassword),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Create Data",
			"error":   result.Error.Error(),
		})
		return
	}

	// Return Status
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Create Data",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	// Get data from req body
	var reqBody struct {
		Username string
		Password string
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Bind Data",
			"error":   err.Error(),
		})
		return
	}

	// Find requested user
	var user models.User
	usernameResult := initializers.DB.First(&user, "username = ?", reqBody.Username)
	if usernameResult.Error != nil || user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Not Found",
		})
		return
	}

	// if user.ID == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "User Not Found",
	// 	})
	// 	return
	// }

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something wrong when creating token...",
		})
		return
	}

	// Send
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"token": tokenString,
	})
}

func GetUserProfile(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
