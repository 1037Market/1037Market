package api

import (
	"1037Market/mysqlDb"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {

		// user identity verify
		cookie := c.Query("user")

		db, err := mysqlDb.GetConnection()
		defer db.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such user")
			return
		}

		var userId string
		err = rows.Scan(&userId)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("scan error: %s", err.Error()))
			return
		}

		// check if target exists
		toId := c.Query("toStudentId")
		rows, err = db.Query("select userId from USERS where userId = ?", toId)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, fmt.Sprintf("user: %s not exists", toId))
			return
		}

		content := c.Query("content")

		txn, err := db.Begin()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		defer func() {
			err = txn.Rollback()
			if err != nil {
				log.Println(err)
			}
		}()

		result, err := txn.Exec("insert into COMMENTS(publisherId, receiverId, content) values(?, ?, ?)",
			userId, toId, content)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected < 1 {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		txn.Commit()
		c.String(http.StatusOK, "OK")
	}
}

func QueryCommentList() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Query("studentId")
		db, err := mysqlDb.GetConnection()
		defer db.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}

		rows, err := db.Query("select commentId from COMMENTS where receiverId = ?", userId)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}

		list := make([]int, 0)
		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("scan error: %s", err.Error()))
				return
			}
			list = append(list, id)
		}

		c.JSON(http.StatusOK, list)
	}
}

func GetCommentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentId := c.Query("commentId")

		db, err := mysqlDb.GetConnection()
		defer db.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		rows, err := db.Query("select publisherId, content from COMMENTS where commentId = ?", commentId)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such commentId")
			return
		}

		type Comment struct {
			FromId  string `json:"fromId"`
			Content string `json:"content"`
		}
		var comment Comment
		err = rows.Scan(&comment.FromId, &comment.Content)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("scan error: %s", err.Error()))
			return
		}

		c.JSON(http.StatusOK, comment)
	}
}
