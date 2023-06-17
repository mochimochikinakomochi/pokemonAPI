package pokemon_api

import (
	"encoding/json"
	"net/http"
	"go-server/handlesql"

	_ "github.com/go-sql-driver/mysql"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("pokemonID")
	pokemon := handlesql.SelectPokemon(param)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	json.NewEncoder(w).Encode(pokemon)
}