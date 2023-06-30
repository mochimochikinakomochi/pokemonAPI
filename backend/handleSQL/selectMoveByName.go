package handleSQL

import (
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func SelectMoveByName(moveName string) Move {
	fmt.Println("moveName:", moveName)
	
	db := ConnectDB()
	defer db.Close()

	// Statusが指定されている場合
	query := `
		select moves.moveID, moveName, moveClass, moveType, movePP, moveAccuracy, movePriority, movePower, moveDescription
			from moves
			join move_description on move_description.moveID = moves.moveID
			where moves.moveName = ?`

	rows, err := db.Query(query, &moveName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var move Move
	for rows.Next() {
		var move_id int
		var move_name string
		var move_class string
		var move_type string
		var move_pp int
		var move_accuracy int
		var move_priority int
		var move_power int
		var move_description string

		err := rows.Scan(&move_id, &move_name, &move_class, &move_type, &move_pp, &move_accuracy, &move_priority, &move_power, &move_description)
		if err != nil {
			panic(err)
		}
		move = Move {
			ID:    move_id,
			Name:  move_name,
			Class: move_class,
			Type:  move_type,
			PP: move_pp,
			Accuracy: move_accuracy,
			Priority: move_priority,
			Power: move_power,
			Description: move_description,
		}
	}

	return move
}