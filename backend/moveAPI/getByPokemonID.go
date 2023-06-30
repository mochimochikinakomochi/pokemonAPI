package moveAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)


func GetByPokemonID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonID := query.Get("pokemonID")

	var moves []handleSQL.Move
	moves = handleSQL.SelectMovesByPokemonID(pokemonID)

	err := json.NewEncoder(w).Encode(moves)
	if err != nil {
		return
	} 
}