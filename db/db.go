package db

import (
	"database/sql"
	"log"

	"github.com/Puneet56/planner/planner_orm"
)

var db *planner_orm.Queries

func GetQuery() *planner_orm.Queries {

	if db != nil {
		return db
	}

	database, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/planner")

	if err != nil {
		log.Fatal(err)
	}

	db = planner_orm.New(database)

	return db
}
