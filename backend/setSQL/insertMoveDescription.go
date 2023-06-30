package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MoveDescription struct {
	ID int
	Text string
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

	jsonFile := "move_data_description_900.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}


	// Pokemonのスライスにデコードする
	var moves []MoveDescription
	err = json.Unmarshal(jsonData, &moves)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}

	for _, move := range moves {
		// moves
		// fmt.Println(move.Text)
		insertMove := "INSERT IGNORE INTO move_description (moveID, moveDescription) values (?, ?)"
		_, err = db.Exec(insertMove, move.ID, move.Text)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", move.ID, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}