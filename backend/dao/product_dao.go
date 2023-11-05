package dao

import (
	"1037Market/mysqlDb"
)

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
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
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
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
		}
		lst = append(lst, id)
	}
	return lst, nil
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
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
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
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
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
			return nil, NewErrorDao(ErrTypeDatabaseScanRows, err.Error())
		}
		lst = append(lst, name)
	}
	return lst, nil
}
