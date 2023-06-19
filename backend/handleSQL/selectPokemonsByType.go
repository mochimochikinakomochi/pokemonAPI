package handleSQL

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func SelectPokemonsByType(_type string) []Pokemon {
	fmt.Printf("Type:%s\n", _type)
	db := ConnectDB()
	defer db.Close()

	query := `
	select pokemon_name.pokemonID, pokemon_name.pokemonName, group_concat(type_name.typeName order by type_name.typeID) as types, concat(pokemon_status.H, ",", pokemon_status.A, ",",  pokemon_status.B, ",",  pokemon_status.C, ",",  pokemon_status.D, ",",  pokemon_status.S)
		from pokemon_name
		join pokemon_type on pokemon_name.pokemonID = pokemon_type.pokemonID
		join type_name on type_name.typeID = pokemon_type.typeID
		join pokemon_status on pokemon_status.pokemonID = pokemon_name.pokemonID
		where pokemon_name.pokemonID in (select pokemon_type.pokemonID 
								from pokemon_type	
								join type_name on type_name.typeID = pokemon_type.typeID
								where type_name.typeName = ?)
		group by pokemon_name.pokemonID, pokemon_name.pokemonName`
	
	rows, err := db.Query(query, &_type)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pokemons []Pokemon
	for rows.Next() {
		var pokemon_id int
		var pokemon_name string
		var pokemon_type string
		var pokemon_stats string

		err := rows.Scan(&pokemon_id, &pokemon_name, &pokemon_type, &pokemon_stats)
		if err != nil {
			panic(err)
		}
		pokemons = append(pokemons, Pokemon{
			ID:    pokemon_id,
			Name:  pokemon_name,
			Type:  pokemon_type,
			Stats: pokemon_stats,
		})
	}

	return pokemons
}