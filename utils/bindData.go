package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindData(obj interface{}, c *gin.Context) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		ReturnResponse(http.StatusBadRequest, "failed to bind data", "error", err.Error(), c)
		return false
	}
	return true
}