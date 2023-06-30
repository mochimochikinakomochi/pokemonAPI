package moveAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

/* type Move struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Class string `json:"Stats"`
	Type  string `json:"Type"`
	PP int `json:"PP"`
} */

func GetByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	moveName := query.Get("moveName")

	move := handleSQL.SelectMoveByName(moveName)

	err := json.NewEncoder(w).Encode(move)
	if err != nil {
		return
	} 
}