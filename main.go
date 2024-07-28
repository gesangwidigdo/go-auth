package main

import (
	"example.com/authentication/controllers"
	"example.com/authentication/initializers"
	"example.com/authentication/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/profile", middlewares.RequireAuth, controllers.GetUserProfile)
	r.Run()
}
