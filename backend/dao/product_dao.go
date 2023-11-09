package dao

import (
	"1037Market/ds"
	"1037Market/mysqlDb"
	"math/rand"
	"strconv"
	"time"
)

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

func PublishProduct(userId string, product ds.ProductPublished) (int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return -1, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		return -1, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer txn.Rollback()
	rand.Seed(time.Now().UnixNano())
	productId := rand.Intn(100000000)
	result, err := txn.Exec("insert into PRODUCTS(productId, userId, title, price, description, createTime, updateTime, status) "+
		"values(?, ?, ? ,?, ?, ?, ?, ?)", productId, userId, product.Title, product.Price, product.Content, time.Now(), time.Now(), "common")
	if err != nil {
		return productId, NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return productId, NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return productId, NewErrorDao(ErrTypeDatabaseExec, userId+"insert product failed")
	}

	for _, uri := range product.ImageURIs {
		result, err = txn.Exec("insert into PRODUCT_IMAGES values(?, ?)", productId, uri)
		if err != nil {
			return productId, NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return productId, NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return productId, NewErrorDao(ErrTypeDatabaseExec, userId+"insert product's images failed")
		}
	}

	for _, category := range product.Categories {
		result, err = txn.Exec("insert into PRODUCT_CATEGORIES values(?, ?)", productId, category)
		if err != nil {
			return productId, NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return productId, NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return productId, NewErrorDao(ErrTypeDatabaseExec, userId+"insert product's categories failed")
		}
	}

	txn.Commit()
	return productId, nil
}

func GetProductById(productId string) (ds.ProductGot, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return ds.ProductGot{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from PRODUCTS where productId = ?", productId)
	if err != nil {
		return ds.ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return ds.ProductGot{}, NewErrorDao(ErrTypeNoSuchProduct, productId+" not found")
	}

	var product ds.ProductGot
	err = rows.Scan(&product.ProductId, &product.Publisher, &product.Title, &product.Price, &product.Status, &product.Content,
		&product.PublishTime, &product.UpdateTime)
	if err != nil {
		return ds.ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
	}

	rows.Close()
	rows, err = db.Query("select imagePath from PRODUCT_IMAGES where productId = ?", productId)
	if err != nil {
		return ds.ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

	product.ImageURIs = make([]string, 0)
	for rows.Next() {
		var uri string
		err = rows.Scan(&uri)
		if err != nil {
			return ds.ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		product.ImageURIs = append(product.ImageURIs, uri)
	}

	rows.Close()
	rows, err = db.Query("select category from PRODUCT_CATEGORIES where productId = ?", productId)
	if err != nil {
		return ds.ProductGot{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

	product.Categories = make([]string, 0)
	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		if err != nil {
			return ds.ProductGot{}, NewErrorDao(ErrTypeScanRows, err.Error())
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

func DeleteProduct(userId string, productId int) error {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	result, err := db.Exec("delete from PRODUCTS where userId = ? and productId = ?", userId, productId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, strconv.Itoa(productId)+" never published")
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

func GetRecommendProductList(seed string, startIdx string, cnt string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	seedInt, err := strconv.ParseInt(seed, 10, 64)
	if err != nil {
		return nil, NewErrorDao(ErrTypeIntParse, err.Error())
	}

	startIdxInt, err := strconv.ParseInt(startIdx, 10, 64)
	if err != nil {
		return nil, NewErrorDao(ErrTypeIntParse, err.Error())
	}

	cntInt, err := strconv.ParseInt(cnt, 10, 64)
	if err != nil {
		return nil, NewErrorDao(ErrTypeIntParse, err.Error())
	}

	rows, err := db.Query("select productId from PRODUCTS")
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	lst := make([]int, 0)
	for rows.Next() {
		var id int
		if err = rows.Scan(&id); err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, id)
	}

	if startIdxInt >= int64(len(lst)) || len(lst) == 0 {
		return make([]int, 0), nil
	}

	rand.Seed(seedInt)
	rand.Shuffle(len(lst), func(i, j int) {
		lst[i], lst[j] = lst[j], lst[i]
	})

	if startIdxInt+cntInt <= int64(len(lst)) {
		return lst[startIdxInt : startIdxInt+cntInt], nil
	} else {
		return lst[startIdxInt:], nil
	}
}

func UpdateProduct(userId string, product ds.ProductUpdated) error {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()
	txn, err := db.Begin()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer txn.Rollback()

	result, err := txn.Exec("update PRODUCTS set title = ?, price = ?, status = ?, description = ?, updateTime = ? where productId = ? and userId = ?",
		product.Title, product.Price, product.Status, product.Content, time.Now(), product.ProductId, userId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, userId+"update product failed")
	}

	result, err = txn.Exec("delete from PRODUCT_CATEGORIES where productId = ?", product.ProductId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err = result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, userId+"update product's categories failed")
	}
	for _, category := range product.Categories {
		result, err = txn.Exec("insert into PRODUCT_CATEGORIES values(?, ?)", product.ProductId, category)
		if err != nil {
			return NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return NewErrorDao(ErrTypeNoSuchProduct, userId+"update product's categories failed")
		}
	}
	result, err = txn.Exec("delete from PRODUCT_IMAGES where productId = ?", product.ProductId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err = result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, userId+"update product's images failed")
	}
	for _, imageURI := range product.ImageURIs {
		result, err = txn.Exec("insert into PRODUCT_IMAGES values(?, ?)", product.ProductId, imageURI)
		if err != nil {
			return NewErrorDao(ErrTypeDatabaseExec, err.Error())
		}
		affected, err = result.RowsAffected()
		if err != nil {
			return NewErrorDao(ErrTypeAffectRows, err.Error())
		}
		if affected < 1 {
			return NewErrorDao(ErrTypeNoSuchProduct, userId+"update product's images failed")
		}
	}
	txn.Commit()
	return nil
}
