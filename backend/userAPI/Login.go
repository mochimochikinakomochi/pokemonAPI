package userAPI

import (
	"SQL/handleSQL"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("kuegffqghwiwofhyhfwhfowihyo3yplqko")

func Login(w http.ResponseWriter, r *http.Request) {
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

	userName := r.FormValue("userName")
	password := r.FormValue("password")
	//fmt.Printf("userName\n", userName)
	//fmt.Printf("password\n", password)

	correctPassword, err := handleSQL.SelectUserPasswordByUserName(userName)
	if err != nil {
		fmt.Printf("error:", err)
	}

	if password == correctPassword {
		tokenString, err := createToken(userName)
		if err != nil {
			log.Fatal(err)
		}

		if tokenString == "" {
			http.Error(w, "トークンの生成に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
		log.Println("Login::Token:", tokenString)
	} else {
		http.Error(w, "Invalid userName or password", http.StatusUnauthorized)
	}
}


func createToken(userName string) (string, error) {
	claims := jwt.MapClaims{
		"userName": userName,
		"exp":		time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}