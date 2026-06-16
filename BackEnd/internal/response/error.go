package response

import "github.com/gin-gonic/gin"

func Error(
	c *gin.Context,
	statusCode int,
	message string,
	error error,
) {
	var errorMessage any
	if error != nil {
		errorMessage = error.Error()
	}
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   errorMessage,
		Data:    nil,
	})
}
