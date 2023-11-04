package dao

import (
	"1037Market/mysqlDb"
	"errors"
)

type Subscribe struct {
	UserId    string
	ProductId string
}

func InsertSubscribe(subscribe Subscribe) error {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	txn, err := db.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	result, err := txn.Exec("insert into SUBSCRIBES(userId, productId) values(?, ?)",
		subscribe.UserId, subscribe.ProductId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("insert failed")
	}
	txn.Commit()
	return nil
}

func GetSubscribes(userId string) ([]int, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select productId from SUBSCRIBES where userId = ?", userId)
	if err != nil {
		return nil, err
	}
	lst := make([]int, 0)
	for rows.Next() {
		var productId int
		if err = rows.Scan(&productId); err != nil {
			return nil, err
		}
		lst = append(lst, productId)
	}
	return lst, nil
}

func DeleteSubscribe(subscribe Subscribe) error {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	txn, err := db.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	result, err := txn.Exec("delete from SUBSCRIBES where userId = ? and productId = ?",
		subscribe.UserId, subscribe.ProductId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("delete failed")
	}
	txn.Commit()
	return nil
}
