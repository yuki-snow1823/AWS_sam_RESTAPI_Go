package mydb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@(db:3306)/go-academy")
}