package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonID := query.Get("pokemonID")

	pokemons := handleSQL.SelectPokemonByID(pokemonID)

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}