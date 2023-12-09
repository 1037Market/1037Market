package middleware

import (
	"1037Market/dao"
	"1037Market/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.GetInstance()

func UserCookieCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		log.Info("user cookie: ", cookie)
		_, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			c.Abort()
		}
		c.Next()
	}
}

func RequestValidationCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info(c.Request.Method, c.Request.Header, c.Request.RequestURI, c.Request.Body)

		c.Next()
	}
}
