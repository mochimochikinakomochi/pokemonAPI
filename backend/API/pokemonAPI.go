package API

import (
	"encoding/json"
	"net/http"
	"./handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("pokemonID")
	pokemon := handleSQL.selectPokemon(param)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	json.NewEncoder(w).Encode(pokemon)
}