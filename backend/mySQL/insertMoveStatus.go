package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Move struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Type string `json:"types"`
	Class string `json:"class"`
	Power int `json:"power"`
	PP int `json:"PP"`
	Accuracy int `json:"Accuracy"`
	Priority int `json:"Priority"`
}

type Class struct {
	ID int
	Name string
}

func main() {
	dbHost := "localhost"
	dbPort := 3306
	dbUser := "root"
	dbPassword := "root"
	dbName := "mysql-pokemonDB"

	jsonFile := "move_data_900.json"

	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err.Error())
	}

	// Pokemonのスライスにデコードする
	var moves []Move
	err = json.Unmarshal(jsonData, &moves)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %s", err.Error())
	}
	

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %s", err.Error())
	}
	defer db.Close()

	/* // move_class_name
	classes := []Class{
		{ID: 1, Name: "physical"},
		{ID: 2, Name: "special"},
		{ID: 3, Name: "status"},
	}
	
	for _, class := range classes {
		insertQuery := "insert ignore into move_class_name (moveClassID, moveClassName) values (?, ?)"
		_, err = db.Exec(insertQuery, class.ID, class.Name)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", class.Name, err.Error())
			} 
	} */


	for _, move := range moves {
		// move_name
		/* insertQuery := "insert ignore into move_name (moveID, moveName) values (?, ?)"
		_, err := db.Exec(insertQuery, move.ID, move.Name)
		if err != nil {
			log.Printf("Failed to insert data for move %s: %s", move.Name, err.Error())
		} */

		/* // move_type
		insertQuery = "INSERT IGNORE INTO move_type (moveID, moveTypeID) VALUES (?, (select type_name.typeID from type_name where ? = type_name.typeName))"
			_, err = db.Exec(insertQuery, move.ID, move.Type)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} */

		// move_class 
		/* insertQuery = "insert ignore into move_class (moveID, moveClassID) values (?, (select moveClassID from move_class_name where ? = move_class_name.moveClassName))"
		_, err = db.Exec(insertQuery, move.ID, move.Class)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} */

		/* // move_power 
		insertQuery = "insert ignore into move_power (moveID, movePower) values (?, ?)"
		_, err = db.Exec(insertQuery, move.ID, move.Power)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} */

		/* // move_PP 
		insertQuery = "insert ignore into move_PP (moveID, movePP) values (?, ?)"
		_, err = db.Exec(insertQuery, move.ID, move.PP)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} */

		/// move_Accuracy 
		insertQuery := "insert ignore into move_accuracy (moveID, moveAccuracy) values (?, ?)"
		_, err = db.Exec(insertQuery, move.ID, move.Accuracy)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} 

		/* // move_Priority
		insertQuery = "insert ignore into move_priority (moveID, movePriority) values (?, ?)"
		_, err = db.Exec(insertQuery, move.ID, move.Priority)
			if err != nil {
				log.Printf("Failed to insert data for Pokemon %s: %s", move.Name, err.Error())
			} */
	}

	log.Println("データの挿入が完了しました。")
}