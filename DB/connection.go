package mydb

import

func Connection() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@(db:3306)/go-academy")
}