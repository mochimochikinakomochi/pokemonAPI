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
	ID     int      `json:"id"`
	Stats []int `json:"stats"`
}


func iPS() {
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

	jsonFile := "pokemon_data_1010_stats.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}

	// Pokemonのスライスにデコードする
	var pokemons []Pokemon
	err = json.Unmarshal([]byte(jsonData), &pokemons)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}


	for _, pokemon := range pokemons {
		// pokemons
		insertPokemon_name := "INSERT IGNORE INTO pokemon_status (pokemonID, H, A, B, C, D, S) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = db.Exec(insertPokemon_name, pokemon.ID, pokemon.Stats[0], pokemon.Stats[1], pokemon.Stats[2], pokemon.Stats[3], pokemon.Stats[4], pokemon.Stats[5])
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.ID, err.Error())
		}
	}

	log.Println("データの挿入が完了しました。")
}