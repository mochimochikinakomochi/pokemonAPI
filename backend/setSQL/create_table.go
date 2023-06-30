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
		`create table if not exists types (
			typeID int not null primary key,
			typeNameEn varchar(255) not null,
			typeNameJa varchar(255) not null
		)`,
		`create table if not exists pokemons (
			pokemonID int primary key,
			pokemonName varchar(255) not null unique,
			pokemonCategory varchar(255) not null,
			pokemonHeight int not null,
			pokemonWeight int not null,
			pokemonImageURL varchar(255) not null
		)`,
		`create table if not exists pokemon_type (
			pokemonID int not null references pokemons(pokemonID),
			pokemonType varchar(255) not null references types(typeNameJa),
			primary key (pokemonID, pokemonType)
		)`,
		`create table if not exists pokemon_status (
			pokemonID int not null references pokemons(pokemonID),
			H int not null,
			A int not null,
			B int not null,
			C int not null,
			D int not null,
			S int not null,
			primary key (pokemonID)
		)`,
		`create table if not exists pokemon_move (
			pokemonID int not null references pokemons(pokemonID),
			moveID int not null references moves(moveID)
		)`,
		`create table if not exists move_class (
			moveClassID int not null primary key,
			moveClassNameEn varchar(255) not null,
			moveClassNameJa varchar(255) not null
		)`,
		`create table if not exists moves (
			moveID int not null unique primary key,
			moveName varchar(255) not null unique,
			moveType varchar(255) not null references types(typeNameJa),
			moveClass varchar(255) not null references move_class(moveClassJa),
			movePP int not null,
			moveAccuracy int not null,
			movePriority int not null,
			movePower int not null
		)`,
		`create table if not exists move_description (
			moveID int not null unique references moves(moveID),
			moveDescription varchar(255) not null,
			primary key (moveID)
		)`,
	}

	for _, query := range createTableQueries {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Failed to create table: %s, %s", query, err.Error())
			return
		}
	}
}