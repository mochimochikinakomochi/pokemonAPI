package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	ImageURL  string `json:"imageURL"`
}

func main() {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "mysql-pokemonDB"

	jsonFile := "pokemon_data_1010_imageURL.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}

	// Pokemonのスライスにデコードする
	var pokemons []Pokemon
	err = json.Unmarshal(jsonData, &pokemons)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}
	

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}
	defer db.Close()

	for _, pokemon := range pokemons {
		insertQuery := "insert ignore into pokemon_image (pokemonID, imageURL) values (?, ?)"
		_, err := db.Exec(insertQuery, pokemon.ID, pokemon.ImageURL)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}