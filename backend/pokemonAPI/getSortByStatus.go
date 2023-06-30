package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonsSortedStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	orderByStatus := query.Get("Status")

	pokemons := handleSQL.SelectAllPokemons(orderByStatus)

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}