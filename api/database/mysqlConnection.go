package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

func Mysql() *sql.DB {

	conf := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("MARIADB_ROOT_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "localhost:3308",
		DBName:               os.Getenv("MARIADB_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}

func DeferClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
