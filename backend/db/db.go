package db

import (
	"database/sql"
	"time"
)

type DataBase struct {
	db *sql.DB
}

func NewDB() *DataBase {
	db, err := sql.Open("mysql", "root:uniquefranky@/market")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DataBase{
		db: db,
	}
}

func (d *DataBase) SetMaxOpenConns(n int) {
	d.db.SetMaxOpenConns(n)
}

func (d *DataBase) SetIdleConns(n int) {
	d.db.SetMaxIdleConns(n)
}

func (d *DataBase) SetConnMaxLifeTime(time2 time.Duration) {
	d.db.SetConnMaxLifetime(time2)
}
