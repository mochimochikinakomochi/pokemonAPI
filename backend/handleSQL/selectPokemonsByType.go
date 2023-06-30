package handleSQL

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)


func SelectPokemonsByType(_type string, status string) []Pokemon {
	fmt.Println("Type:%", _type, "Status:%", status)

	db := ConnectDB()
	defer db.Close()

	var query string
	var rows *sql.Rows
	var err error

	// Statusが指定されている場合
	if status == "N" {
		query = `
			select pokemons.pokemonID, pokemons.pokemonName, pokemons.pokemonCategory, pokemons.pokemonHeight, pokemons.pokemonWeight, pokemons.pokemonImageURL, group_concat(pokemon_type.pokemonType) as pokemonTypes, concat(pokemon_status.H, ",", pokemon_status.A, ",",  pokemon_status.B, ",",  pokemon_status.C, ",",  pokemon_status.D, ",",  pokemon_status.S) as pokemonStatus
				from pokemons
				join pokemon_type on pokemon_type.pokemonID = pokemons.pokemonID
				join pokemon_status on pokemon_status.pokemonID = pokemons.pokemonID
				where pokemons.pokemonID in (select pokemonID 
												from pokemon_type	
												where pokemonType = ?)
				group by pokemons.pokemonID
				order by pokemons.pokemonID`
		
		rows, err = db.Query(query, &_type)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	} else if status == "SUM" {
		query = `
			select pokemons.pokemonID, pokemons.pokemonName, pokemons.pokemonCategory, pokemons.pokemonHeight, pokemons.pokemonWeight, pokemons.pokemonImageURL, group_concat(pokemon_type.pokemonType) as pokemonTypes, concat(pokemon_status.H, ",", pokemon_status.A, ",",  pokemon_status.B, ",",  pokemon_status.C, ",",  pokemon_status.D, ",",  pokemon_status.S) as pokemonStatus
				from pokemons
				join pokemon_type on pokemon_type.pokemonID = pokemons.pokemonID
				join pokemon_status on pokemon_status.pokemonID = pokemons.pokemonID
				where pokemons.pokemonID in (select pokemonID 
												from pokemon_type	
												where pokemonType = ?)
				group by pokemons.pokemonID
				order by (pokemon_status.H + pokemon_status.A + pokemon_status.B + pokemon_status.C + pokemon_status.D + pokemon_status.S) desc`
		
		rows, err = db.Query(query, &_type)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	} else {
		query = `
			select pokemons.pokemonID, pokemons.pokemonName, pokemons.pokemonCategory, pokemons.pokemonHeight, pokemons.pokemonWeight, pokemons.pokemonImageURL, group_concat(pokemon_type.pokemonType) as pokemonTypes, concat(pokemon_status.H, ",", pokemon_status.A, ",",  pokemon_status.B, ",",  pokemon_status.C, ",",  pokemon_status.D, ",",  pokemon_status.S) as pokemonStatus
				from pokemons
				join pokemon_type on pokemon_type.pokemonID = pokemons.pokemonID
				join pokemon_status on pokemon_status.pokemonID = pokemons.pokemonID
				where pokemons.pokemonID in (select pokemonID 
												from pokemon_type	
												where pokemonType = ?)
				group by pokemons.pokemonID`
		query = fmt.Sprintf("%s order by pokemon_status.%s desc", query, status)
		rows, err = db.Query(query, &_type)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
	}
	

	var pokemons []Pokemon
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