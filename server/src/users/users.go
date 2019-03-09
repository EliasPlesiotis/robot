package users

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Password string
	Email string
}

func CreateUser() error {
	db, err := sql.Open("mysql", "root:moyskeman54@tcp(127.0.0.1:3306)/newsblog")
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}