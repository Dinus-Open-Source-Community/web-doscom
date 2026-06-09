package response

import "github.com/gin-gonic/gin"

func Error(
	c *gin.Context,
	statusCode int,
	message string,
	error any,
) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   error,
		Data:    nil,
	})
}
