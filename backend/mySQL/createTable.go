package mySQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createTable() {
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

	createTableQueries := []string {
		`create table if not exists user_password (
			userName varchar(255) primary key,
			userPassword varchar(255)
		)`,
		`create table if not exists user_favPokemonID (
			userName varchar(255) not null references user_password(userName),
			userFavPokemonID int not null references pokemon_name(pokemonID),
			primary key (userName, userFavPokemonID)
		)`,
		`create table if not exists pokemon_name (
			pokemonID int primary key,
			pokemonName varchar(255) not null unique
		)`,
		`create table if not exists type_name (
			typeID int primary key,
			typeName varchar(255) not null unique
		)`,
		`create table if not exists pokemon_type (
			pokemonID int not null references pokemon_name(pokemonID),
			typeID varchar(255) not null references type_name(typeID),
			primary key (pokemonID, typeID)
		)`,
		`create table if not exists pokemon_status (
			pokemonID int not null references pokemon_name(pokemonID),
			H int not null,
			A int not null,
			B int not null,
			C int not null,
			D int not null,
			S int not null,
			primary key (pokemonID)
		)`,
		`create table if not exists pokemon_height (
			pokemonID int not null references pokemon_name(pokemonID),
			height int not null
		)`,
		`create table if not exists pokemon_weight (
			pokemonID int not null references pokemon_name(pokemonID),
			weight int not null
		)`,
	}

	for _, query := range createTableQueries {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Failed to create table: %s", err.Error())
			return
		}
	}
}