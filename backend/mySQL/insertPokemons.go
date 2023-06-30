package mySQL

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
	Types  []string `json:"types"`
	Weight int      `json:"weight"`
	Height int      `json:"height"`
}

type _Type struct {
	ID int 
	Name string
}

func insertPokemons() {
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

	jsonFile := "pokemon_data_1010.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}

	types := []_Type {
		{ID: 1, Name: "normal"},
		{ID: 2, Name: "fire"},
		{ID: 3, Name: "water"},
		{ID: 4, Name: "electric"},
		{ID: 5, Name: "grass"},
		{ID: 6, Name: "ice"},
		{ID: 7, Name: "fighting"},
		{ID: 8, Name: "poison"},
		{ID: 9, Name: "ground"},
		{ID: 10, Name: "flying"},
		{ID: 11, Name: "psychic"},
		{ID: 12, Name: "bug"},
		{ID: 13, Name: "rock"},
		{ID: 14, Name: "ghost"},
		{ID: 15, Name: "dragon"},
		{ID: 16, Name: "dark"},
		{ID: 17, Name: "steel"},
		{ID: 18, Name: "fairy"}}

		typesJa := []_Type {
			{ID: 1, Name: "ノーマル"},
			{ID: 2, Name: "ほのお"},
			{ID: 3, Name: "みず"},
			{ID: 4, Name: "でんき"},
			{ID: 5, Name: "くさ"},
			{ID: 6, Name: "こおり"},
			{ID: 7, Name: "かくとう"},
			{ID: 8, Name: "どく"},
			{ID: 9, Name: "じめん"},
			{ID: 10, Name: "ひこう"},
			{ID: 11, Name: "エスパー"},
			{ID: 12, Name: "むし"},
			{ID: 13, Name: "いわ"},
			{ID: 14, Name: "ゴースト"},
			{ID: 15, Name: "ドラゴン"},
			{ID: 16, Name: "あく"},
			{ID: 17, Name: "はがね"},
			{ID: 18, Name: "フェアリー"}}

	// type_name
	for _, _type := range typesJa {
		insertType_name := "INSERT IGNORE INTO type_name (typeID, typeName) VALUES (?, ?)"
		_, err = db.Exec(insertType_name, _type.ID, _type.Name)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", _type.Name, err.Error())
		}
	}

	// Pokemonのスライスにデコードする
	var pokemons []Pokemon
	err = json.Unmarshal(jsonData, &pokemons)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}

	for _, pokemon := range pokemons {
		/* // pokemon_name
		insertPokemon_name := "INSERT IGNORE INTO pokemon_name (pokemonID, pokemonName) VALUES (?, ?)"
		_, err = db.Exec(insertPokemon_name, pokemon.ID, pokemon.Name)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
		} */

		// pokemon_type
		for _, _type := range pokemon.Types {
			insertPokemon_type := "INSERT IGNORE INTO pokemon_type (pokemonID, typeID) VALUES (?, (select type_name.typeID from type_name where ? = type_name.typeName))"
			_, err = db.Exec(insertPokemon_type, pokemon.ID, _type)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
			}
		}

		/* // pokemon_height
		insertPokemon_height := "INSERT IGNORE INTO pokemon_height (pokemonID, height) VALUES (?, ?)"
		_, err = db.Exec(insertPokemon_height, pokemon.ID, pokemon.Height)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
		} */

		/* // pokemon_weight
		insertPokemon_weight := "INSERT IGNORE INTO pokemon_weight (pokemonID, weight) VALUES (?, ?)"
		_, err = db.Exec(insertPokemon_weight, pokemon.ID, pokemon.Weight)
		if err != nil {
			log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.Name, err.Error())
		} */
	}

	log.Println("データの挿入が完了しました。")
}