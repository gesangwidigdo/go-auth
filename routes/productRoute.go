package routes

import (
	"example.com/authentication/controllers"
	"example.com/authentication/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.RouterGroup) {
	r.POST("/", middlewares.RequireAuth, controllers.CreateProduct)
	r.GET("/", middlewares.RequireAuth, controllers.GetAllProduct)
	r.GET("/:id", middlewares.RequireAuth, controllers.GetProductByID)
	r.PUT("/:id", middlewares.RequireAuth, controllers.UpdateProduct)
	r.DELETE("/:id", middlewares.RequireAuth, controllers.DeleteProduct)
}
