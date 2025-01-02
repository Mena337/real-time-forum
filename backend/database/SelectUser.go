package database

func GetUser(email string) (string, error) {
	query := `SELECT password  FROM users WHERE email = ?`
	var password string
	err := db.QueryRow(query, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
