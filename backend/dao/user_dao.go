package dao

import (
	"1037Market/ds"
	"1037Market/mysqlDb"
	"errors"
)

func AddNewUser(user ds.RegisterUser) error {
	// get database connection
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return err
	}

	//start transaction
	txn, err := db.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	// insert new user
	result, err := txn.Exec("insert into USERS values(?, ?)", user.StudentId, user.HashedPsw)
	if err != nil {
		return err
	}

	// check result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return errors.New("用户已存在")
	}

	// insert into USER_INFOS
	result, err = txn.Exec("insert into USER_INFOS(userId, nickName, avatar, contact) values(?, ?, ?, ?)",
		user.StudentId, user.StudentId, "null", "null")
	if err != nil {
		return err
	}
	// check result
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("用户已存在")
	}

	txn.Commit()
	return nil
}

func GetUserIdByCookie(cookie string) (string, error) {
	db, err := mysqlDb.GetConnection()
	defer db.Close()
	if err != nil {
		return "", err
	}

	rows, err := db.Query("select userId from COOKIES where cookie = ?", cookie)
	if err != nil {
		return "", err
	}

	if !rows.Next() {
		return "", errors.New("invalid cookie")
	}
	var userId string
	if err = rows.Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}
