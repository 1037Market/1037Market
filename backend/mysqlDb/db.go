package mysqlDb

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

type DataBase struct {
	*sql.DB
}

func newDB() (*DataBase, error) {
	file, err := os.Open("/var/MYSQLPASSWORD")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	// 创建一个Scanner用于读取文件
	scanner := bufio.NewScanner(file)

	// 读取第一行
	var psw string
	if scanner.Scan() {
		psw = scanner.Text() // Text方法返回不带换行符的当前行
	}

	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	db, err := sql.Open("mysql", "root:"+psw+"@/market")
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

func GetConnection() (*DataBase, error) {
	return newDB()
}
