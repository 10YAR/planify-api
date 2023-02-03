package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

func Mysql() *sql.DB {

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}

	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		DB_USER = "root"
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		DB_PASSWORD = "root"
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		DB_PORT = "3308"
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		DB_NAME = "planify"
	}

	conf := mysql.Config{
		User:                 DB_USER,
		Passwd:               DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 DB_HOST + ":" + DB_PORT,
		DBName:               DB_NAME,
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
