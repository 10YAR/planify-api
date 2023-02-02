package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"os"
)

func mysqlConnection() {

	conf := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("MARIADB_ROOT_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "db:3306",
		DBName:               os.Getenv("MARIADB_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
}
