package handleSQL

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)


func SelectPokemonByName(pokemonName string) Pokemon {
	fmt.Printf("pokemonName:%s\n", pokemonName)
	
	db := ConnectDB()
	defer db.Close()

	query := `
		select pokemons.pokemonID, pokemons.pokemonName, pokemons.pokemonCategory, pokemons.pokemonHeight, pokemons.pokemonWeight, pokemons.pokemonImageURL, group_concat(pokemon_type.pokemonType) as pokemonTypes, concat(pokemon_status.H, ",", pokemon_status.A, ",",  pokemon_status.B, ",",  pokemon_status.C, ",",  pokemon_status.D, ",",  pokemon_status.S) as pokemonStatus
			from pokemons
			join pokemon_type on pokemon_type.pokemonID = pokemons.pokemonID
			join pokemon_status on pokemon_status.pokemonID = pokemons.pokemonID
			where pokemons.pokemonName = ?
			group by pokemons.pokemonID`
	
	rows, err := db.Query(query, &pokemonName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pokemon Pokemon
	for rows.Next() {
		var pokemon_id int
		var pokemon_name string
		var pokemon_type string
		var pokemon_stats string
		var pokemon_category string
		var pokemon_height int
		var pokmeon_weight int
		var pokemon_imageURL string

		err := rows.Scan(&pokemon_id, &pokemon_name, &pokemon_category, &pokemon_height, &pokmeon_weight, &pokemon_imageURL, &pokemon_type, &pokemon_stats)
		if err != nil {
			panic(err)
		}

		pokemon_types := strings.Split(pokemon_type, ",")

		statusArray := strings.Split(pokemon_stats, ",")
		h, _ := strconv.Atoi(statusArray[0])
		a, _ := strconv.Atoi(statusArray[1])
		b, _ := strconv.Atoi(statusArray[2])
		c, _ := strconv.Atoi(statusArray[3])
		d, _ := strconv.Atoi(statusArray[4])
		s, _ := strconv.Atoi(statusArray[5])
		pokemon_status := Pokemon_status {
			H: h,
			A: a,
			B: b,
			C: c,
			D: d,
			S: s,
		}

		pokemon = Pokemon {
			ID:    pokemon_id,
			Name:  pokemon_name,
			Types:  pokemon_types,
			Status: pokemon_status,
			Category: pokemon_category,
			Height: pokemon_height,
			Weight: pokmeon_weight,
			ImageURL: pokemon_imageURL,
		}
	}

	return pokemon
}