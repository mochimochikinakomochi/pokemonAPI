package moveAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)


func GetByClass(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	class := query.Get("Class")

	var moves []handleSQL.Move
	moves = handleSQL.SelectMovesByClass(class)

	err := json.NewEncoder(w).Encode(moves)
	if err != nil {
		return
	} 
}