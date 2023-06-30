package handleSQL

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func SelectUserPasswordByUserName(userName string) (string, error) {
	db := ConnectDB()
	defer db.Close()

	query := `
		select userPassword from user_password where userName = ?`
	
	rows, err := db.Query(query, &userName)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var userPassword string
	found := false
	for rows.Next() {
		err := rows.Scan(&userPassword)
		if err != nil {
			return "", err
		}
		found = true
	}

	if !found {
		return "", errors.New("User not found")
	}

	return userPassword, nil
}