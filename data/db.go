package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func (product *Products) DbCreateObject() (*sql.DB , error) {
	db, err := sql.Open("mysql","root:@/bookstore")
	if err != nil {
		return nil , err
	}
	return db , err
}
