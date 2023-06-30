package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonsByMoveID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	moveID := query.Get("moveID")

	pokemons := handleSQL.SelectPokemonsByMoveID(moveID)

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}