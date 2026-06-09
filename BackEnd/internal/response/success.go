package response

import "github.com/gin-gonic/gin"

func Success(
	c *gin.Context,
	message string,
	statusCode int,
	data any,
) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
