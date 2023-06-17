package package2

import (
	"encoding/json"
	"net/http"
	"handleSQL/package1"

	_ "github.com/go-sql-driver/mysql"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("pokemonID")
	pokemon := package1.SelectPokemon(param)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	json.NewEncoder(w).Encode(pokemon)
}