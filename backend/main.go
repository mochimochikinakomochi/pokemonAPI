package main

import (
	"SQL/pokemonAPI"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", pokemonAPI.HandleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
