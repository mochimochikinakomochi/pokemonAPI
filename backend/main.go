package main

import (
	"fmt"
	"log"
	"net/http"
	"SQL/pokemonAPI"
)

func main() {
	http.HandleFunc("/", pokemonAPI.HandleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}