package api

import (
	"1037Market/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSubscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		productId := c.Query("productId")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}

		userId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if err = dao.InsertSubscribe(dao.Subscribe{UserId: userId, ProductId: productId}); err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}

		c.String(http.StatusOK, "ok")
	}
}

func GetSubscribes() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("userId")
		lst, err := dao.GetSubscribes(userId)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}

func DeleteSubscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		productId := c.Query("productId")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}

		userId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if err = dao.DeleteSubscribe(dao.Subscribe{UserId: userId, ProductId: productId}); err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}

		c.String(http.StatusOK, "ok")
	}
}
