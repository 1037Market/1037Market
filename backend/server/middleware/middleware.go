package middleware

import (
	"1037Market/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.GetInstance()

func UserCookieCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		log.Info("user cookie: ", cookie)
		// 这里放用户的cookie检查（看有没有这个cookie）
		// 因为前端接口还没更新，这里什么也不做
		c.Next()
	}
}

func RequestValidationCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(c.Request.Method, c.Request.Header, c.Request.RequestURI, c.Request.Body)

		c.Next()
	}
}
