package api

import (
	"1037Market/dao"
	"1037Market/mysqlDb"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func PublishProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Product struct {
			Title      string   `json:"title"`
			Content    string   `json:"content"`
			Categories []string `json:"categories"`
			ImageURIs  []string `json:"imageURIs"`
			Price      float32  `json:"price"`
		}

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
		defer rows.Close()
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such user")
			return
		}
		var userId string
		err = rows.Scan(&userId)
		if err != nil {
			c.String(http.StatusInternalServerError, "scan error")
			return
		}

		// get product info
		var product Product
		err = c.ShouldBindJSON(&product)
		if err != nil {
			c.String(http.StatusBadRequest, "incorrect request format")
			return
		}

		txn, err := db.Begin()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		defer func() {
			if err := txn.Rollback(); err != nil {
				log.Println(err)
			}
		}()

		rand.Seed(time.Now().UnixNano())
		productId := rand.Intn(100000000)
		result, err := txn.Exec("insert into PRODUCTS(productId, userId, title, price, description, createTime, updateTime, status) values(?, ?, ? ,?, ?, ?, ?, ?)", productId, userId,
			product.Title, product.Price, product.Content, time.Now(), time.Now(), "common")
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected < 1 {
			c.String(http.StatusInternalServerError, "please try again")
			return
		}

		for _, uri := range product.ImageURIs {
			result, err = txn.Exec("insert into PRODUCT_IMAGES values(?, ?)", productId, uri)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
				return
			}
			if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected < 1 {
				c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
				return
			}
		}

		for _, category := range product.Categories {
			result, err = txn.Exec("insert into PRODUCT_CATEGORIES values(?, ?)", productId, category)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
				return
			}
			if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected < 1 {
				c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("%d", productId))
		if err = txn.Commit(); err != nil {
			log.Println(err)
		}
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
			// c.String(http.StatusInternalServerError, "database error: %s", err.Error())
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
			// c.String(http.StatusInternalServerError, "database error: %s", err.Error())
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
			// c.String(http.StatusInternalServerError, "database error: %s", err)
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
			// c.String(http.StatusInternalServerError, "database error: %s", err)
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
			// c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}
