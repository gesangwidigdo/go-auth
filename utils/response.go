package utils

import "github.com/gin-gonic/gin"

func ReturnResponse(StatusCode int, msg string, key string, data interface{}, c *gin.Context) {
	if (key == "" && data == nil) {
		c.JSON(StatusCode, gin.H{
			"message": msg,
		})
	} else {
		c.JSON(StatusCode, gin.H{
			"message": msg,
			key: data,
		})
	}
}