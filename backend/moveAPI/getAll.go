package moveAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)


func GetAll(w http.ResponseWriter, r *http.Request) {
	var moves []handleSQL.Move
	moves = handleSQL.SelectAllMoves()

	err := json.NewEncoder(w).Encode(moves)
	if err != nil {
		return
	} 
}