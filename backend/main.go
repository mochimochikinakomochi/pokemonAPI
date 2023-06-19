package main

import (
	"fmt"
	"log"
	"net/http"
	"SQL/pokemonAPI"
	"SQL/userAPI"
)

func main() {
	http.HandleFunc("/", pokemonAPI.HandleAPI)
	http.HandleFunc("/insert/", userAPI.HandleUserApi)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}