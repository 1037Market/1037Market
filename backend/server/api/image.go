package api

import (
	"1037Market/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		// identity verify
		cookie := c.Query("user")
		_, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			handleError(c, err)
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		imagePrefix := dao.GenerateRandomDigits(5)
		dst := "./uploads/" + imagePrefix + file.Filename
		if err = c.SaveUploadedFile(file, dst); err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeSysSaveFile, err.Error()))
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
			handleError(c, dao.NewErrorDao(dao.ErrTypeSysOpenFile, err.Error()))
			return
		}
		c.File(filePath)
	}
}
