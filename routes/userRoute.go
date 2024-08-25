package routes

import (
	"example.com/authentication/controllers"
	"example.com/authentication/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/profile", middlewares.RequireAuth, controllers.GetUserProfile)
}
