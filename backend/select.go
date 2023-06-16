package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func mai() {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "mysql-db"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}
	defer db.Close()

	query := `
	select pokemons.id, pokemons.name, group_concat(type_master.type_name order by type_master.id) as types, CONCAT(stats.H, ',', stats.A, ',', stats.B, ',', stats.C, ',', stats.D, ',', stats.S) as stats
		from pokemons
		join types on types.pokemonID = pokemons.id
		join type_master on type_master.id = types.typeID
		join stats on stats.pokemonID = pokemons.id
		group by pokemons.id, pokemons.name
		order by pokemons.id`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()


	for rows.Next() {
		var pokemons_id int
		var pokemons_name string
		var pokemon_type string
		var pokemon_stats string
		err := rows.Scan(&pokemons_id, &pokemons_name, &pokemon_type, &pokemon_stats)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", pokemons_id, "NAME:", pokemons_name, "pokemon_type", pokemon_type, "pokemon_stats", pokemon_stats)

	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

}


