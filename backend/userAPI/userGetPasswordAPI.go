package userAPI

import (
	"SQL/handleSQL"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)


func HandleUserGetPasswordApi(w http.ResponseWriter, r *http.Request) {
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
	userName := query.Get("userName")

	correctPassword, err := handleSQL.SelectUserPasswordByUserName(userName)
	if err != nil {
		fmt.Printf("error:", err)
	} else {
		fmt.Println("user password:", correctPassword)
	}

	err = json.NewEncoder(w).Encode(correctPassword)
	if err != nil {
		return
	}
}