package api

import (
	"1037Market/dao"
	"1037Market/ds"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		cookie := c.Query("user")
		userId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			handleError(c, err)
			return
		}
		productId := c.Query("productId")

		product, err := dao.GetProductById(userId, productId)

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
		signString := c.Query("sign")
		sign, err := strconv.ParseBool(signString)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		lst, err := dao.GetProductListByKeyword(keyword, sign)
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
		userId, err := dao.GetUserIdByCookie(cookie)
		if err != nil {
			handleError(c, err)
			return
		}
		productIdString := c.Query("productId")
		productId, err := strconv.Atoi(productIdString)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeIntParse, err.Error()))
			return
		}

		err = dao.DeleteProduct(userId, productId)
		if err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func GetRandomProductList() gin.HandlerFunc { // Deprecated
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

func GetRecommendProductList() gin.HandlerFunc {
	return func(c *gin.Context) {
		seed := c.Query("seed")
		startIndex := c.Query("startIndex")
		count := c.Query("count")
		signStr := c.Query("sign")
		sign, err := strconv.ParseBool(signStr)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}

		lst, err := dao.GetRecommendProductList(seed, startIndex, count, sign)
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
		startIndexString := c.Query("startIndex")
		signStr := c.Query("sign")
		sign, err := strconv.ParseBool(signStr)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeWrongRequestFormat, err.Error()))
			return
		}
		startIndex, err := strconv.Atoi(startIndexString)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeIntParse, err.Error()))
			return
		}
		cnt := c.Query("count")
		lst, err := dao.GetProductListByCategory(category, startIndex, cnt, sign)
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

func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		userId, err := dao.UserIdentityVerify(cookie)
		if err != nil {
			handleError(c, err)
			return
		}

		var product ds.ProductUpdated
		if err = c.ShouldBindJSON(&product); err != nil {
			c.String(http.StatusBadRequest, "incorrect request format")
			return
		}
		if err = dao.UpdateProduct(userId, product); err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func SoldProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Query("user")
		userId, err := dao.UserIdentityVerify(cookie)
		if err != nil {
			handleError(c, err)
			return
		}

		productIdString := c.Query("productId")
		productId, err := strconv.Atoi(productIdString)
		if err != nil {
			handleError(c, dao.NewErrorDao(dao.ErrTypeIntParse, err.Error()))
			return
		}

		if err := dao.SoldProduct(userId, productId); err != nil {
			handleError(c, err)
			return
		}
		c.String(http.StatusOK, "OK")
	}
}
