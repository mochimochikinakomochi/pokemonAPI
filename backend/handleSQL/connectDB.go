package handleSQL

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "mysql-pokemonDB"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}

	return db
}