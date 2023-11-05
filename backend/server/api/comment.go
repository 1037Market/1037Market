package api

import (
	"1037Market/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		fromId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			handleError(c, err)
			return
		}
		toId := c.Query("toStudentId")
		content := c.Query("content")
		starsString := c.Query("stars")
		stars, err := strconv.Atoi(starsString) // 转换字符串为整型
		if err != nil || stars < 0 || stars > 5 {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		err = dao.CreateComment(fromId, toId, content, stars)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func QueryCommentList() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("studentId")
		list, err := dao.QueryCommentList(userId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, list)
	}
}

func GetCommentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId := c.Query("commentId")
		comment, err := dao.GetCommentById(commentId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, comment)
	}
}
