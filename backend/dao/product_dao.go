package dao

import (
	"1037Market/mysqlDb"
)

func GetProductListByCategory(category, count string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("select productId from PRODUCT_CATEGORIES where category = ? order by rand() limit ?",
		category, count)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	lst := make([]int, 0)
	for rows.Next() {
		var productId int
		err = rows.Scan(&productId)
		if err != nil {
			return nil, err
		}
		lst = append(lst, productId)
	}
	return lst, nil
}

func GetCategoryList() ([]string, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select categoryName from CATEGORIES")
	if err != nil {
		return nil, err
	}

	lst := make([]string, 0)
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, err
		}
		lst = append(lst, name)
	}

	return lst, nil
}
