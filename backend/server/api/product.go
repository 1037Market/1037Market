package api

import (
	"1037Market/dao"
	"1037Market/ds"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		userId, err := dao.UserIdentityVerify(cookie)
		if err != nil {
			handleError(c, err)
			return
		}

		var product ds.ProductPublished
		err = c.ShouldBindJSON(&product)
		if err != nil {
			c.String(http.StatusBadRequest, "incorrect request format")
			return
		}
		productId, err := dao.PublishProduct(userId, product)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%d", productId))
	}
}

func GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("productId")

		product, err := dao.GetProductById(id)

		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func GetProductListByKeyword() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("keyword")

		lst, err := dao.GetProductListByKeyword(keyword)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}

func GetProductListByStudentId() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentId := c.Query("studentId")

		lst, err := dao.GetProductListByStudentId(studentId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		productId := c.Query("productId")

		err := dao.DeleteProduct(cookie, productId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func GetRandomProductList() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnt := c.Query("count")
		lst, err := dao.GetRandomProductList(cnt)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}

func GetProductListByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		cnt := c.Query("count")
		lst, err := dao.GetProductListByCategory(category, cnt)
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}

func GetCategoryList() gin.HandlerFunc {
	return func(c *gin.Context) {
		lst, err := dao.GetCategoryList()
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}
