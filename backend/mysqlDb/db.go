package mysqlDb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DataBase struct {
	*sql.DB
}

func newDB() (*DataBase, error) {
	db, err := sql.Open("mysql", "root:uniquefranky@/market")
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DataBase{
		db,
	}, nil
}

var d *DataBase

func GetDefault() (*DataBase, error) {
	if nil == d {
		var err error
		d, err = newDB()
		if err != nil {
			return nil, err
		}
	}
	return d, nil

}
