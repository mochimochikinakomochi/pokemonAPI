package myAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("pokemonID")
	pokemon := handleSQL.SelectPokemon(param)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	
	err := json.NewEncoder(w).Encode(pokemon)
	if err != nil {
		return
	}
}