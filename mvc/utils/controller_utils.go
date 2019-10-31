package utils

import (
	"github.com/gin-gonic/gin"
)

// Respond returns either a JSON or XML payload containing the response
func Respond(c *gin.Context, status int, body interface{}){
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

// RespondError returns either a JSON or XML payload containing information about an error
func RespondError(c *gin.Context, err *ApplicationError){
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}