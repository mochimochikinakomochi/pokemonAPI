package moveAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)


func GetByType(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	moveType := query.Get("Type")

	moves := handleSQL.SelectMovesByType(moveType)

	err := json.NewEncoder(w).Encode(moves)
	if err != nil {
		return
	} 
}