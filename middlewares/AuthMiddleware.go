package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"example.com/authentication/initializers"
	"example.com/authentication/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get Header
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is missing",
		})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token format",
		})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	tokenString := authToken[1]
	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or expired token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token Expired",
			})
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		// Find user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}
}
