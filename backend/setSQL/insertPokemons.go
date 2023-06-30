package setSQL

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	ID     int      `json:"ID"`
	Name   string   `json:"Name"`
	Category string `json:"Category"`
	Height int      `json:"Height"`
	Weight int      `json:"Weight"`
	ImageURL string `json:"ImageURL"`
}

func iP() {
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

	jsonFile := "pokemons_1010.json"

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

	for _, pokemon := range pokemons {
		// pokemons
		insertPokemon_name := "INSERT IGNORE INTO pokemons (pokemonID, pokemonName, pokemonCategory, pokemonHeight, pokemonWeight, pokemonImageURL) VALUES (?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(insertPokemon_name, pokemon.ID, pokemon.Name, pokemon.Category, pokemon.Height, pokemon.Weight, pokemon.ImageURL)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}