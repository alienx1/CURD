package send

import (
	"github.com/gin-gonic/gin"
)

func SendModel(c *gin.Context, data any, err error, status int) {
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"status": "error",
			"error":  err,
			"data":   data,
		})
	} else {
		c.JSON(status, gin.H{
			"status": "pass",
			"error":  err,
			"data":   data,
		})
	}
}
