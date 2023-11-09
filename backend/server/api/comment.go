package api

import (
	"1037Market/dao"
	"1037Market/ds"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		fromId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			handleError(c, err)
			return
		}
		var comment ds.CommentSent
		if err := c.ShouldBindJSON(&comment); err != nil || comment.Stars < 0 || comment.Stars > 5 {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		commentId, err := dao.CreateComment(fromId, comment)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "%d", commentId)
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
