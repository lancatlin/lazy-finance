package main

import "github.com/gin-gonic/gin"

func errorHandler(c *gin.Context) {
	c.Next() // process request

	// Check if there's an error
	if len(c.Errors) > 0 {
		// Use the last error as it's the most recent one
		err := c.Errors.Last().Err
		c.JSON(c.Writer.Status(), gin.H{"message": err.Error()})
	}
}
