package handleSQL

import (
	_ "github.com/go-sql-driver/mysql"
)

type Pokemon struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Stats string `json:"Stats"`
}

func SelectPokemon(pokemonID string) Pokemon {
	db := ConnectDB()

	query := `
	select pokemons.id, pokemons.name, group_concat(type_master.type_name order by type_master.id) as types, CONCAT(stats.H, ',', stats.A, ',', stats.B, ',', stats.C, ',', stats.D, ',', stats.S) as stats
		from pokemons
		join types on types.pokemonID = pokemons.id
		join type_master on type_master.id = types.typeID
		join stats on stats.pokemonID = pokemons.id
		where pokemons.id = ?
		group by pokemons.id, pokemons.name
		order by pokemons.id`

	rows, err := db.Query(query, &pokemonID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var pokemon Pokemon
	for rows.Next() {
		var pokemon_id int
		var pokemon_name string
		var pokemon_type string
		var pokemon_stats string

		err := rows.Scan(&pokemon_id, &pokemon_name, &pokemon_type, &pokemon_stats)
		if err != nil {
			panic(err)
		}
		pokemon = Pokemon{
			ID:    pokemon_id,
			Name:  pokemon_name,
			Type:  pokemon_type,
			Stats: pokemon_stats,
		}
	}

	return pokemon
}
