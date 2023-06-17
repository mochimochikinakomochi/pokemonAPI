package main 

import (
	"fmt"
	"log"
	"net/http"
	"myAPI/package2"
)

func main() {
	http.HandleFunc("/", package2.HandleAPI)
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}