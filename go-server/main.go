package main 

import (
	"fmt"
	"log"
	"net/http"
	"go-server/pokemon_api"
)

func main() {
	http.HandleFunc("/", pokemon_api.HandleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}