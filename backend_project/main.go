package main 

import (
	"fmt"
	"log"
	"net/http"
	"SQL/myAPI"
)

func main() {
	http.HandleFunc("/", myAPI.HandleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}