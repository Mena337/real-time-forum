package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"RTF/backend/database"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req Log
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Println(req)

	password, err := database.GetUser(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	if password != req.Password {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(password)

	// resp := Log{
	// 	Email: req.Email,
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(resp)
}
