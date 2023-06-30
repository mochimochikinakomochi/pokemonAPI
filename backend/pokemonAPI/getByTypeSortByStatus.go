package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonsByTypeAndSortedStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pokemonType := query.Get("Type")
	orderByStatus := query.Get("Status")

	pokemons := handleSQL.SelectPokemonsByType(pokemonType, orderByStatus)

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}