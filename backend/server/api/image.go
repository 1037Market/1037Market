package api

import (
	"1037Market/mysqlDb"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {

		// identity verify
		cookie, err := c.Cookie("user")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}
		db, err := mysqlDb.GetNewDb()
		defer db.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error")
		}
		rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		defer rows.Close()
		if !rows.Next() { // no corresponding cookie is stored
			c.String(400, "用户不存在")
			return
		}

		// get the image
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		// save the image
		imagePrefix := generateRandomDigits(5)
		dst := fmt.Sprintf("./uploads/%s%s", imagePrefix, file.Filename)
		if err = c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("%s%s", imagePrefix, file.Filename))
	}
}

func DownloadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Query("imageURI")
		filePath := path.Join("./uploads", uri)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.String(http.StatusBadRequest, fmt.Sprintf("image %s not exists", uri))
		}
		c.File(filePath)
	}
}
