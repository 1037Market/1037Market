package mysqlDb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"time"
)

type DataBase struct {
	*sql.DB
}

func newDB() (*DataBase, error) {
	psw, err := ioutil.ReadFile("/var/MYSQLPASSWORD")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	db, err := sql.Open("mysql", "root:"+string(psw)+"@/market")
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

func GetNewDb() (*DataBase, error) {
	return newDB()
}
