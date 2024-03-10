package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/planner")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MySQL database")
}

func GetDB() *sql.DB {
	return db
}
