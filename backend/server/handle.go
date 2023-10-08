package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to 1037 Market!")
	}
}
