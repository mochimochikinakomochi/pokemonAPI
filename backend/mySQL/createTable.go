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
		`create table if not exists pokemons (
			pokemonID int primary key,
			pokemonName varchar(255) not null unique,
			pokemonHeight int not null,
			pokemonWeight int not null,
			pokemonImageURL varchar(255) not null
		)`,
		/* `create table if not exists pokemon_name (
			pokemonID int primary key,
			pokemonName varchar(255) not null unique
		)`, */
		`create table if not exists type_name (
			typeID int primary key,
			typeName varchar(255) not null unique
		)`,
		`create table if not exists pokemon_type (
			pokemonID int not null references pokemons(pokemonID),
			pokemonType varchar(255) not null references type_name(typeName),
			primary key (pokemonID, typeID)
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
		/* `create table if not exists pokemon_height (
			pokemonID int not null references pokemon_name(pokemonID),
			height int not null
		)`, */
		/* `create table if not exists pokemon_weight (
			pokemonID int not null references pokemon_name(pokemonID),
			weight int not null
		)`, */
		/* `create table if not exists pokemon_image (
			pokemonID int not null references pokemon_name(pokemonID),
			imageURL varchar(255) not null,
			primary key (pokemonID)
		)`, */
		`create table if not exists move_name (
			moveID int not null primary key,
			moveName varchar(255) not null unique
		)`,
		`create table if not exists move_type (
			moveID int not null,
			moveTypeID int not null,
			primary key (moveID),
			foreign key (moveID) references move_name(moveID),
			foreign key (moveTypeID) references type_name(typeID) 
		)`,
		`create table if not exists move_class_name (
			moveClassID int not null primary key,
			moveClassName varchar(255) not null
		)`,
		`create table if not exists move_class (
			moveID int not null ,
			moveClassID int not null,
			primary key (moveID),
			foreign key (moveID) references move_name(moveID),
			foreign key (moveClassID) references move_class_name(moveClassID)
		)`,
		`create table if not exists move_power (
			moveID int not null primary key,
			movePower int not null,
			primary key (moveID)
		)`,
		`create table if not exists move_pp (
			moveID int not null,
			movePP int not null,
			primary key (moveID),
			foreign key (moveID) references move_name(moveID)
		)`,
		`create table if not exists move_accuracy (
			moveID int not null,
			moveAccuracy int not null,
			primary key (moveID),
			foreign key (moveID) references move_name(moveID)
		)`,
		`create table if not exists move_priority (
			moveID int not null,
			movePriority int not null,
			primary key (moveID),
			foreign key (moveID) references move_name(moveID)
		)`,
		`create table if not exists pokemon_category (
			pokemonID int not null,
			pokemonCategory varchar(255) not null,
			primary key (pokemonID),
			foreign key (pokemonID) references pokemon_name(pokemonID)
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