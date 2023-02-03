package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func Mysql() *sql.DB {

	conf := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3308",
		DBName:               "planify",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}

func DoQuery(query string, params ...any) (*sql.Rows, error) {
	db := Mysql()
	res, err := db.Query(query, params...)
	if err != nil {
		return res, err
	}
	DeferClose(db)
	return res, nil
}

func DeferClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
