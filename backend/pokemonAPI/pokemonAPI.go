package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonID := query.Get("ID")
	pokemonType := query.Get("Type")
	var pokemons []handleSQL.Pokemon
	if pokemonID != "" {
		pokemons = handleSQL.SelectPokemonByID(pokemonID)
	} else if pokemonType != ""{
		pokemons = handleSQL.SelectPokemonsByType(pokemonType)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}