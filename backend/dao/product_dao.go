package dao

import (
	"1037Market/mysqlDb"
	"log"
	"math/rand"
	"time"
)

type ProductGot struct {
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

type ProductPublished struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
	ImageURIs  []string `json:"imageURIs"`
	Price      float32  `json:"price"`
}

func UserIdentityVerify(cookie string) (string, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
	if err != nil {
		return "", NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	if !rows.Next() {
		return "", NewErrorDao(ErrTypeNoSuchUser, "cookie "+cookie+" not found")
	}
	var userId string
	err = rows.Scan(&userId)
	if err != nil {
		return "", NewErrorDao(ErrTypeScanRows, err.Error())
	}
	return userId, nil
}

func PublishProduct(userId string, product ProductPublished) error {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer func() {
		if err := txn.Rollback(); err != nil {
			log.Println(err)
		}
	}()
	rand.Seed(time.Now().UnixNano())
	productId := rand.Intn(100000000)
	result, err := txn.Exec("insert into PRODUCTS(productId, userId, title, price, description, createTime, updateTime, status) "+
		"values(?, ?, ? ,?, ?, ?, ?, ?)", productId, userId, product.Title, product.Price, product.Content, time.Now(), time.Now(), "common")
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeDatabaseExec, userId+"insert product failed")
	}

	for _, uri := range product.ImageURIs {
		result, err = txn.Exec("insert into PRODUCT_IMAGES values(?, ?)", productId, uri)
		if err != nil {
			return NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return NewErrorDao(ErrTypeDatabaseExec, userId+"insert product's images failed")
		}
	}

	for _, category := range product.Categories {
		result, err = txn.Exec("insert into PRODUCT_CATEGORIES values(?, ?)", productId, category)
		if err != nil {
			return NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return NewErrorDao(ErrTypeDatabaseExec, userId+"insert product's categories failed")
		}
	}

	if err = txn.Commit(); err != nil {
		log.Println(err)
	}
	return nil
}

func GetProductById(productId string) (ProductGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return ProductGot{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from PRODUCTS where productId = ?", productId)
	if err != nil {
		return ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return ProductGot{}, NewErrorDao(ErrTypeNoSuchProduct, productId+" not found")
	}

	var product ProductGot
	err = rows.Scan(&product.ProductId, &product.Publisher, &product.Title, &product.Price, &product.Status, &product.Content,
		&product.PublishTime, &product.UpdateTime)
	if err != nil {
		return ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
	}

	rows.Close()
	rows, err = db.Query("select imagePath from PRODUCT_IMAGES where productId = ?", productId)
	if err != nil {
		return ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	product.ImageURIs = make([]string, 0)
	for rows.Next() {
		var uri string
		err = rows.Scan(&uri)
		if err != nil {
			return ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		product.ImageURIs = append(product.ImageURIs, uri)
	}

	rows.Close()
	rows, err = db.Query("select category from PRODUCT_CATEGORIES where productId = ?", productId)
	if err != nil {
		return ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	product.Categories = make([]string, 0)
	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		if err != nil {
			return ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		product.Categories = append(product.Categories, category)
	}
	return product, nil
}

func GetProductListByKeyword(keyword string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select productId from PRODUCTS where title like ? or description like ?",
		"%"+keyword+"%", "%"+keyword+"%")
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, id)
	}
	return lst, nil
}

func GetProductListByStudentId(studentId string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select productId from PRODUCTS where userId = ?", studentId)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, id)
	}
	return lst, nil
}

func DeleteProduct(cookie, productId string) error {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	if !rows.Next() {
		return NewErrorDao(ErrTypeNoSuchUser, err.Error())
	}

	var userId string
	if err = rows.Scan(&userId); err != nil {
		return NewErrorDao(ErrTypeScanRows, err.Error())
	}
	defer rows.Close()

	result, err := db.Exec("delete from PRODUCTS where userId = ? and productId = ?", userId, productId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, productId+" never published")
	}
	return nil
}

func GetRandomProductList(count string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select productId from PRODUCTS order by rand() limit ?", count)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var productId int
		err = rows.Scan(&productId)
		if err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, productId)
	}
	return lst, nil
}

func GetProductListByCategory(category, count string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select productId from PRODUCT_CATEGORIES where category = ? order by rand() limit ?",
		category, count)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var productId int
		err = rows.Scan(&productId)
		if err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, productId)
	}
	return lst, nil
}

func GetCategoryList() ([]string, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select categoryName from CATEGORIES")
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]string, 0)
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, name)
	}
	return lst, nil
}