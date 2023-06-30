package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonName := query.Get("pokemonName")

	pokemon := handleSQL.SelectPokemonByName(pokemonName)

	err := json.NewEncoder(w).Encode(pokemon)
	if err != nil {
		return
	} 
}