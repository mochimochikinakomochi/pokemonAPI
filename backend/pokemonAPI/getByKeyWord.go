package pokemonAPI

import (
	"encoding/json"
	"net/http"
	"SQL/handleSQL"

	_ "github.com/go-sql-driver/mysql"
)

func GetPokemonsByKeyWord(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	keyWord := query.Get("KeyWord")

	pokemons := handleSQL.SelectPokemonsByKeyWord(keyWord)

	err := json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	} 
}