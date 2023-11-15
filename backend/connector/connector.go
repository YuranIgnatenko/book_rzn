package connector

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "ecommerce"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func Connect() {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
		return
	}
	db.Ping()
}
