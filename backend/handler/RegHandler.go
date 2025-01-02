package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"RTF/backend/database"
)

func RegHnadler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	return
	// }
	// db, err := sql.Open("sqlite3", "./forum.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	var req Reg
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// // Validate email
	// if req.Email == "" {
	// 	http.Error(w, "Email is required", http.StatusBadRequest)
	// 	return
	// }

	err = database.RegisterUser(req.Nickname, req.First_Name, req.Last_Name, req.Email, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// resp := Reg{
	// 	Email: userEmail,
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(resp)
}
