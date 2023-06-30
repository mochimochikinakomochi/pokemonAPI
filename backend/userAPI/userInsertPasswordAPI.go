package userAPI

import (
	"net/http"
	"SQL/handleSQL"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"Name"`
	PassWord string `json:"Password"`
}

func HandleUserInsertPasswordApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
        // OPTIONSメソッドへのリクエストに対するCORSヘッダーを設定する
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	query := r.URL.Query()
	fmt.Printf("query:%s", query)
	userName := query.Get("Name")
	password := query.Get("Password")
	fmt.Printf("userName:%s, password: %s", userName, password)

	db := handleSQL.ConnectDB()
	defer db.Close()

	insertQuery := `insert into user_password (userName, userPassword) values (?, ?)`
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Printf("Failed to insert statement: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, password)
	if err != nil {
		log.Printf("Failed to insert data: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}