package setSQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type _Type struct {
	ID int 
	Name string
}

func iT() {
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


	typesEn := []_Type {
		{ID: 1, Name: "normal"},
		{ID: 2, Name: "fire"},
		{ID: 3, Name: "water"},
		{ID: 4, Name: "electric"},
		{ID: 5, Name: "grass"},
		{ID: 6, Name: "ice"},
		{ID: 7, Name: "fighting"},
		{ID: 8, Name: "poison"},
		{ID: 9, Name: "ground"},
		{ID: 10, Name: "flying"},
		{ID: 11, Name: "psychic"},
		{ID: 12, Name: "bug"},
		{ID: 13, Name: "rock"},
		{ID: 14, Name: "ghost"},
		{ID: 15, Name: "dragon"},
		{ID: 16, Name: "dark"},
		{ID: 17, Name: "steel"},
		{ID: 18, Name: "fairy"}}

	typesJa := []_Type {
		{ID: 1, Name: "ノーマル"},
		{ID: 2, Name: "ほのお"},
		{ID: 3, Name: "みず"},
		{ID: 4, Name: "でんき"},
		{ID: 5, Name: "くさ"},
		{ID: 6, Name: "こおり"},
		{ID: 7, Name: "かくとう"},
		{ID: 8, Name: "どく"},
		{ID: 9, Name: "じめん"},
		{ID: 10, Name: "ひこう"},
		{ID: 11, Name: "エスパー"},
		{ID: 12, Name: "むし"},
		{ID: 13, Name: "いわ"},
		{ID: 14, Name: "ゴースト"},
		{ID: 15, Name: "ドラゴン"},
		{ID: 16, Name: "あく"},
		{ID: 17, Name: "はがね"},
		{ID: 18, Name: "フェアリー"}}

		// types
		for i := 0; i < len(typesEn); i++  {
			insertPokemon_type := "INSERT IGNORE INTO types (typeID, typeNameEn, typeNameJa) VALUES (?, ?, ?)"
			_, err = db.Exec(insertPokemon_type, i+1, typesEn[i].Name, typesJa[i].Name)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", i, err.Error())
			}
		}

	log.Println("データの挿入が完了しました。")
}