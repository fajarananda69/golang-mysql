package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1)/whatsapp")
	if err != nil {
		fmt.Println("gagal")
		return nil, err
	}
	fmt.Println("success login")
	return db, nil

}
