package setSQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MoveClass struct {
	ID int
	NameEn string
	NameJa string
}


func iMC() {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "mysql-pokemonDB"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}
	defer db.Close()

	moveClasses := []MoveClass {
		{ID: 1, NameEn: "physical", NameJa: "物理"},
		{ID: 2, NameEn: "special", NameJa: "特殊"},
		{ID: 3, NameEn: "status", NameJa: "変化"},
	}

	for _, moveClass := range moveClasses {
		// moves
		insertMoveClass := "INSERT IGNORE INTO move_class (moveClassID, moveClassNameEn, moveClassNameJa) values (?, ?, ?)"
		_, err = db.Exec(insertMoveClass, moveClass.ID, moveClass.NameEn, moveClass.NameJa)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", moveClass.ID, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}