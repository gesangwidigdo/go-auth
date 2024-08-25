package routes

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	// user route
	userRoutes := r.Group("user")
	UserRoute(userRoutes)

	// product route
	productRoutes := r.Group("product")
	ProductRoute(productRoutes)
}
