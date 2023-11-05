package dao

import (
	"1037Market/mysqlDb"
)

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

func GetProductById(productId string) (Product, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return Product{}, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from PRODUCTS where productId = ?", productId)
	if err != nil {
		return Product{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	var product Product
	if !rows.Next() {
		return Product{}, NewErrorDao(ErrTypeNoSuchProduct, err.Error())
	}
	err = rows.Scan(&product.ProductId, &product.Publisher, &product.Title, &product.Price, &product.Status, &product.Content,
		&product.PublishTime, &product.UpdateTime)
	if err != nil {
		return Product{}, NewErrorDao(ErrTypeScanRows, err.Error())
	}

	rows.Close()
	rows, err = db.Query("select imagePath from PRODUCT_IMAGES where productId = ?", productId)
	if err != nil {
		return Product{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	product.ImageURIs = make([]string, 0)
	for rows.Next() {
		var uri string
		err = rows.Scan(&uri)
		if err != nil {
			return Product{}, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		product.ImageURIs = append(product.ImageURIs, uri)
	}

	rows.Close()
	rows, err = db.Query("select category from PRODUCT_CATEGORIES where productId = ?", productId)
	if err != nil {
		return Product{}, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()

	product.Categories = make([]string, 0)
	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		if err != nil {
			return Product{}, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		product.Categories = append(product.Categories, category)
	}
	return product, nil
}

func GetProductListByKeyword(keyword string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}

	rows, err := db.Query("select productId from PRODUCTS where title like ? or description like ?",
		"%"+keyword+"%", "%"+keyword+"%")
	defer rows.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

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
	defer db.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}

	rows, err := db.Query("select productId from PRODUCTS where userId = ?", studentId)
	defer rows.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

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
	defer db.Close()
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
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
		return NewErrorDao(ErrTypeNoSuchProduct, err.Error())
	}
	return nil
}

func GetRandomProductList(count string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}

	rows, err := db.Query("select productId from PRODUCTS order by rand() limit ?", count)
	defer rows.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

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
	defer db.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}

	rows, err := db.Query("select productId from PRODUCT_CATEGORIES where category = ? order by rand() limit ?",
		category, count)
	defer rows.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

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
	defer db.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}

	rows, err := db.Query("select categoryName from CATEGORIES")
	defer rows.Close()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}

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
