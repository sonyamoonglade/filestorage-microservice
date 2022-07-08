package app_errors

import "github.com/gin-gonic/gin"

type ErrorHandler func(ctx *gin.Context)

func Internal(c *gin.Context) {
	c.AbortWithStatusJSON(500, gin.H{
		"message": "internal server error",
	})
	return
}

func InternalMsg(c *gin.Context, v string) {
	c.AbortWithStatusJSON(500, gin.H{
		"message": v,
	})
}
