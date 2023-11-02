package dao

import "1037Market/mysqlDb"

func GetCategoryList() ([]string, error) {
	db, err := mysqlDb.GetConnection()
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
