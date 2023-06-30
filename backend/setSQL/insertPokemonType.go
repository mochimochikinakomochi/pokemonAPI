package setSQL

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type PokemonType struct {
	ID int
	Types []string
}

type _Type struct {
	ID int 
	Name string
}

func iPT() {
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

	jsonFile := "pokemon_Type_data_1010.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}


	// Pokemonのスライスにデコードする
	var pokemons []PokemonType
	err = json.Unmarshal(jsonData, &pokemons)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}

	for _, pokemon := range pokemons {
		// pokemon_type
		for _, _type := range pokemon.Types {
			insertPokemon_type := "INSERT IGNORE INTO pokemon_type (pokemonID, pokemonType) VALUES (?, (select types.typeNameJa from types where types.typeNameEn = ?))"
			_, err = db.Exec(insertPokemon_type, pokemon.ID, _type)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", pokemon.ID, err.Error())
			}
		}
	}

	log.Println("データの挿入が完了しました。")
}