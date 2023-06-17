package handlesql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	c := mysql.Config{
		DBName: "mysql-db",
		User: "root",
		Passwd: "root",
		Addr: "localhost:3306",
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}
	defer db.Close()

	return db
}