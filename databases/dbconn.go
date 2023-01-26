package databases

import (
	"database/sql"
	"go-todolist/utils"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-todolist")
	utils.PanicIfError(err)
	return db, nil
}
