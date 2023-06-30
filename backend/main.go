package main

import (
	"fmt"
	"log"
	"net/http"
	"SQL/pokemonAPI"
	"SQL/userAPI"
	"SQL/moveAPI"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("kuegffqghwiwofhyhfwhfowihyo3yplqko")

func main() {
	http.HandleFunc("/user/get/token/", userAPI.Login)	

	// pokemonを取得するAPI
	// 全てのポケモンをstatusでsortしてgetするAPI
	// http.Handle("/pokemon/all/", handleCORS(tokenAuthMiddleware(http.HandlerFunc(pokemonAPI.GetPokemonsSortedStatus))))
	http.Handle("/pokemon/all/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonsSortedStatus)))
	// 指定されたタイプを持つポケモンをstatusでsortしてgetするAPI
	// http.Handle("/pokemon/list/type/", handleCORS(tokenAuthMiddleware(http.HandlerFunc(pokemonAPI.GetPokemonsByTypeAndSortedStatus))))
	http.Handle("/pokemon/list/type/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonsByTypeAndSortedStatus)))
	// 指定されたIDのポケモンをgetするAPI
	// http.Handle("/pokemon/byID/", handleCORS(tokenAuthMiddleware(http.HandlerFunc(pokemonAPI.GetPokemonByID))))
	http.Handle("/pokemon/byID/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonByID)))
	// 指定された名前のポケモンをgetするAPI
	// http.Handle("/pokemon/byName/", handleCORS(tokenAuthMiddleware(http.HandlerFunc(pokemonAPI.GetPokemonByName))))
	http.Handle("/pokemon/byName/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonByName)))
	// 指定された技IDを覚えるポケモンをgetするAPI
	http.Handle("/pokemon/byMoveID/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonsByMoveID)))
	// 指定されたkeyWordを名前に含むポケモンをgetするAPI
	http.Handle("/pokemon/byKeyWord/", handleCORS(http.HandlerFunc(pokemonAPI.GetPokemonsByKeyWord)))



	// 全ての技をgetするAPI
	http.Handle("/move/all/", handleCORS(http.HandlerFunc(moveAPI.GetAll)))
	// 指定されたIDの技をgetするAPI
	http.Handle("/move/byID/", handleCORS(http.HandlerFunc(moveAPI.GetByID)))
	//http.Handle("/move/byID/", handleCORS(http.HandlerFunc(moveAPI.GetByID)))
	// 指定された名前の技をgetするAPI
	http.Handle("/move/byName/", handleCORS(http.HandlerFunc(moveAPI.GetByName)))
	// 指定されたポケモンIDが覚える技をgetするAPI
	http.Handle("/move/byPokemonID/", handleCORS(http.HandlerFunc(moveAPI.GetByPokemonID)))

	

	// 

	//http.HandleFunc("/", pokemonAPI.HandleAPI)
	http.HandleFunc("/user/insert/", userAPI.HandleUserInsertPasswordApi)
	http.HandleFunc("/user/get/password/", userAPI.HandleUserGetPasswordApi)
	//http.HandleFunc("/user/get/token/", userAPI.Login)


	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func tokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		fmt.Println("token:::", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("無効なトークンの署名アルゴリズムです")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "認証されてません", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func tokenAuthCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwtToken")
		fmt.Println("cookie::", cookie)
		if err != nil {
			// cookieが存在しない場合は認証エラーとする
			if err == http.ErrNoCookie {
				http.Error(w, "認証されていません", http.StatusUnauthorized)
				return
			}
			// その他のエラーが発生した場合はエラーレスポンスを返す
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// トークンの認証
		tokenString := cookie.Value
		fmt.Println("token:::", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("無効なトークンの署名アルゴリズムです")
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "認証されてません", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if r.Method == "OPTIONS" {
			return 
		}

		next.ServeHTTP(w, r)
	})
}