package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Move struct {
	ID int
	Name string
	Type string
	Class string
	PP int
	Accuracy int
	Priority int
	Power int
}


func main() {
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

	jsonFile := "move_data_900.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}


	// Pokemonのスライスにデコードする
	var moves []Move
	err = json.Unmarshal(jsonData, &moves)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}

	for _, move := range moves {
		// moves
		insertMove := "INSERT IGNORE INTO moves (moveID, moveName, moveType, moveClass, movePP, moveAccuracy, movePriority, movePower) VALUES (?, ?, (select types.typeNameJa from types where types.typeNameEn = ?), (select moveClassNameJa from move_class where moveClassNameEn = ?), ?, ?, ?, ?)"
		_, err = db.Exec(insertMove, move.ID, move.Name, move.Type, move.Class, move.PP, move.Accuracy, move.Priority, move.Power)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", move.ID, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}