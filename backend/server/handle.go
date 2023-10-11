package server

import (
	"1037Market/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func helloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {

		product := ds.ProductInfo{
			ProductId:   114,
			UserID:      514,
			Title:       "title",
			Description: "description",
			Images:      make([]string, 2),
			Price:       123,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}

		c.JSON(http.StatusOK, product)
	}
}
