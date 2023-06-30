package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Pokemon_move struct {
	PokemonID int `json:"pokemonID"`
	MoveID int `json:"moveID"`
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

	jsonFile := "pokemon_move_data_1010.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}



	// Pokemonのスライスにデコードする
	var moves [][]Pokemon_move
	err = json.Unmarshal(jsonData, &moves)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}

	for _, pokemon := range moves {
		for _, move := range pokemon {
			insertMove := "INSERT IGNORE INTO pokemon_move (pokemonID, moveID) VALUES (?, ?)"
			_, err = db.Exec(insertMove, &move.PokemonID, &move.MoveID)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.PokemonID, err.Error())
			}
		}
	}

	log.Println("データの挿入が完了しました。")
}