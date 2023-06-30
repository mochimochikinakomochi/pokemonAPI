package handleSQL

import (
	"fmt"
	"log"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)


func SelectPokemonsByKeyWord(keyWord string) []Pokemon {
	fmt.Printf("word:%s\n", keyWord)
	
	db := ConnectDB()
	defer db.Close()

	keyWord = strings.Join([]string{keyWord, "%"}, "")

	query := `
		select pokemons.pokemonID, pokemons.pokemonName, pokemons.pokemonImageURL
			from pokemons
			where pokemons.pokemonName like ?
			order by pokemons.pokemonID
			`
	
	rows, err := db.Query(query, &keyWord)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pokemons []Pokemon
	for rows.Next() {
		var pokemon_id int
		var pokemon_name string
		pokemon_types := []string{"", ""}
		var pokemon_category string
		var pokemon_height int
		var pokmeon_weight int
		var pokemon_imageURL string

		err := rows.Scan(&pokemon_id, &pokemon_name, &pokemon_imageURL)
		if err != nil {
			panic(err)
		}

		pokemon_status := Pokemon_status {
			H: 0,
			A: 0,
			B: 0,
			C: 0,
			D: 0,
			S: 0,
		}

		pokemons = append(pokemons, Pokemon {
			ID:    pokemon_id,
			Name:  pokemon_name,
			Types:  pokemon_types,
			Status: pokemon_status,
			Category: pokemon_category,
			Height: pokemon_height,
			Weight: pokmeon_weight,
			ImageURL: pokemon_imageURL,
		})
	}

	return pokemons
}