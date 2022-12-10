package errors

import "github.com/gin-gonic/gin"

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Errors) > 0 {
			c.Next()
			err := c.Errors.Last()
			result := gin.H{"error": err.Error()}
			c.AbortWithStatusJSON(ConvertError(err), result)
		}
	}
}
