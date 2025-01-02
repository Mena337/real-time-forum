package database

import "fmt"

func RegisterUser(nickname, first_name, last_name, email, password string) error {
	// Check if db is nil before executing the query
	// if db == nil {
	// 	return fmt.Errorf("database connection is not initialized")
	// }
	query := `INSERT INTO users (nickname, first_name, last_name, email, password) VALUES (?, ?, ?, ?,?)`
	if _, err := db.Exec(query, nickname, first_name, last_name, email, password); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func LoginUser(email, password string) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	_, err := db.Exec(query, email, password)
	return err
}
