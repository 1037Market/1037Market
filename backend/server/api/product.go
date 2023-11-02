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
		cookie, err := c.Cookie("user")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}
		db, err := mysqlDb.GetConnection()
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

		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}

		rows, err := db.Query("select * from PRODUCTS where productId = ?", id)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("database error: %s", err.Error()))
			return
		}
		type Product struct {
			ProductId   int      `json:"productId"`
			Title       string   `json:"title"`
			Content     string   `json:"content"`
			Publisher   string   `json:"publisher"`
			Price       float32  `json:"price"`
			PublishTime string   `json:"publishTime"`
			UpdateTime  string   `json:"updateTime"`
			ImageURIs   []string `json:"imageURIs"`
			Categories  []string `json:"categories"`
			Status      string   `json:"status"`
		}
		var product Product
		if !rows.Next() {
			c.String(http.StatusBadRequest, "no such product")
			return
		}
		err = rows.Scan(&product.ProductId, &product.Publisher, &product.Title, &product.Price, &product.Status, &product.Content,
			&product.PublishTime, &product.UpdateTime)
		if err != nil {
			c.String(http.StatusInternalServerError, "scan error: %s", err.Error())
			return
		}
		rows.Close()
		rows, err = db.Query("select imagePath from PRODUCT_IMAGES where productId = ?", id)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err.Error())
			return
		}

		product.ImageURIs = make([]string, 0)
		for rows.Next() {
			var uri string
			err = rows.Scan(&uri)
			if err != nil {
				c.String(http.StatusInternalServerError, "scan error: %s", err.Error())
				return
			}
			product.ImageURIs = append(product.ImageURIs, uri)
		}
		rows.Close()

		rows, err = db.Query("select category from PRODUCT_CATEGORIES where productId = ?", id)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err.Error())
			return
		}
		product.Categories = make([]string, 0)
		for rows.Next() {
			var category string
			err = rows.Scan(&category)
			if err != nil {
				c.String(http.StatusInternalServerError, "scan error: %s", err.Error())
				return
			}
			product.Categories = append(product.Categories, category)
		}
		rows.Close()

		c.JSON(http.StatusOK, product)
	}
}

func GetProductListByKeyword() gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Query("keyword")

		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err.Error())
			return
		}
		rows, err := db.Query("select productId from PRODUCTS where title like ? or description like ?",
			"%"+keyword+"%", "%"+keyword+"%")
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err.Error())
			return
		}
		rows.Close()
		lst := make([]int, 0)
		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				c.String(http.StatusInternalServerError, "scan error: %s", err.Error())
				return
			}
			lst = append(lst, id)
		}
		c.JSON(http.StatusOK, lst)
	}
}

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		productId := c.Query("productId")
		if err != nil {
			c.String(http.StatusBadRequest, "cookie not set")
			return
		}
		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if !rows.Next() {
			c.String(http.StatusBadRequest, "invalid cookie")
			return
		}

		var userId string
		if err = rows.Scan(&userId); err != nil {
			c.String(http.StatusInternalServerError, "scan error: %s", err)
			return
		}
		rows.Close()

		result, err := db.Exec("delete from PRODUCTS where userId = ? and productId = ?", userId, productId)
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		affected, err := result.RowsAffected()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if affected < 1 {
			c.String(http.StatusBadRequest, "invalid product id")
			return
		}
		c.String(http.StatusOK, "OK")
	}
}

func GetRandomProductList() gin.HandlerFunc {
	return func(c *gin.Context) {
		cnt := c.Query("count")
		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		rows, err := db.Query("select productId from PRODUCTS order by rand() limit ?", cnt)
		defer rows.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		lst := make([]int, 0)
		for rows.Next() {
			var productId int
			if err = rows.Scan(&productId); err != nil {
				c.String(http.StatusInternalServerError, "scan error: %s", err)
				return
			}
			lst = append(lst, productId)
		}
		c.JSON(http.StatusOK, lst)
	}
}

func GetProductListByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		cnt := c.Query("count")
		fmt.Println(category, cnt)
		db, err := mysqlDb.GetConnection()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		rows, err := db.Query("select productId from PRODUCT_CATEGORIES where category = ? order by rand() limit ?",
			category, cnt)
		defer rows.Close()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		lst := make([]int, 0)
		for rows.Next() {
			var productId int
			err = rows.Scan(&productId)
			if err != nil {
				c.String(http.StatusInternalServerError, "scan error: %s", err)
				return
			}
			lst = append(lst, productId)
		}
		c.JSON(http.StatusOK, lst)
	}
}

func GetCategoryList() gin.HandlerFunc {
	return func(c *gin.Context) {
		lst, err := dao.GetCategoryList()
		if err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		c.JSON(http.StatusOK, lst)
	}
}
