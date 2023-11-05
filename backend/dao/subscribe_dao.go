package dao

import (
	"1037Market/mysqlDb"
)

func InsertSubscribe(userId, productId string) error {
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

	result, err := txn.Exec("insert into SUBSCRIBES(userId, productId) values(?, ?)",
		userId, productId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}

	if affected < 1 {
		return NewErrorDao(ErrTypeProductAlreadyExist, "user "+userId+" subscribe "+productId+"already exist")
	}
	txn.Commit()
	return nil
}

func GetSubscribes(userId string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseConnection, err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select productId from SUBSCRIBES where userId = ?", userId)
	if err != nil {
		return nil, NewErrorDao(ErrTypeDatabaseQuery, err.Error())
	}
	defer rows.Close()
	lst := make([]int, 0)
	for rows.Next() {
		var productId int
		if err = rows.Scan(&productId); err != nil {
			return nil, NewErrorDao(ErrTypeScanRows, err.Error())
		}
		lst = append(lst, productId)
	}
	return lst, nil
}

func DeleteSubscribe(userId, productId string) error {
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
	result, err := txn.Exec("delete from SUBSCRIBES where userId = ? and productId = ?",
		userId, productId)
	if err != nil {
		return NewErrorDao(ErrTypeDatabaseExec, err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return NewErrorDao(ErrTypeAffectRows, err.Error())
	}
	if affected < 1 {
		return NewErrorDao(ErrTypeNoSuchProduct, "user "+userId+" delete "+productId+" not exist")
	}
	txn.Commit()
	return nil
}
