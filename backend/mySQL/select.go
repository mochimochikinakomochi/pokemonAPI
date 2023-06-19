package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	query := `
	select pokemon_name.pokemonID, pokemon_name.pokemonName, pokemon_status.H, pokemon_status.A, pokemon_status.B, pokemon_status.C, pokemon_status.D, pokemon_status.S
		from pokemon_name
		join pokemon_status on pokemon_status.pokemonID = pokemon_name.pokemonID
		order by pokemon_name.pokemonID`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pokemons_id int
		var pokemons_name string
		var pokemon_H int
		var pokemon_A int
		var pokemon_B int
		var pokemon_C int
		var pokemon_D int
		var pokemon_S int
		err := rows.Scan(&pokemons_id, &pokemons_name, &pokemon_H, &pokemon_A, &pokemon_B, &pokemon_C, &pokemon_D, &pokemon_S)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", pokemons_id, "NAME:", pokemons_name, "H:", pokemon_H, "A:", pokemon_A, "B:", pokemon_B, "C:", pokemon_C, "D:", pokemon_D, "S:", pokemon_S)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}