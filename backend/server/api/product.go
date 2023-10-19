package api

import (
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
			Price      int      `json:"price"`
		}

		// user identity verify
		cookie, err := c.Cookie("user")
		if err != nil {
			c.String(http.StatusBadRequest, "no cookie is set")
			return
		}
		db, err := mysqlDb.GetNewDb()
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
		result, err := txn.Exec("insert into PRODUCTS values(?, ?, ?, ?, ?, ?)", productId, userId,
			product.Price, product.Content, time.Now(), time.Now())
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
