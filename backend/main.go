package main

import (
	"fmt"
	"log"
	"net/http"
	"backend/API"
)

func main() {
	http.HandleFunc("/", pokemonAPI.handleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}