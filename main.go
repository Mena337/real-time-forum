package main

import (
	"fmt"
	"log"
	"net/http"

	"RTF/backend/database"
	"RTF/backend/handler"
)

func main() {
	// db, err := sql.Open("sqlite3", "./users.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// log.Println("Initializing the database...")
	// db, err := database.InitDB("forum.db")
	// if err != nil {
	// 	log.Fatal("Error initializing the database: ", err)
	// }

	database.InitDB("RTF.db")

	http.Handle("/", http.FileServer(http.Dir("frontend")))

	http.HandleFunc("POST /signup", handler.RegHnadler)
	http.HandleFunc("POST /login", handler.LoginHandler)

	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
