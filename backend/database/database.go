package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dbPath string) error {
	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	db = database
	CreateTables(db)
	// database.InsertDummyData(db)
	log.Println("Database setup complete")
	return nil
}

func CreateTables(db *sql.DB) {
	userTable := `
        CREATE TABLE IF NOT EXISTS users (
            uid INTEGER PRIMARY KEY AUTOINCREMENT,
            nickname TEXT NOT NULL UNIQUE,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            password TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE
        );`

	categoryTable := `
        CREATE TABLE IF NOT EXISTS categories (
            category_id INTEGER PRIMARY KEY AUTOINCREMENT,
            category_name TEXT NOT NULL UNIQUE
        );`

	commentTable := `
        CREATE TABLE IF NOT EXISTS comments (
            comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
            comment TEXT NOT NULL,
            user_id INTEGER,
            post_id INTEGER,
            FOREIGN KEY (user_id) REFERENCES users (uid),
            FOREIGN KEY (post_id) REFERENCES posts (post_id)
        );`

	likeCommentTable := `
        CREATE TABLE IF NOT EXISTS likeComment (
            comment_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (comment_id) REFERENCES comments(comment_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (comment_id, user_id)
        );`

	dislikeCommentTable := `
        CREATE TABLE IF NOT EXISTS dislikeComment (
            comment_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (comment_id) REFERENCES comments(comment_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (comment_id, user_id)
        );`

	likesTable := `
        CREATE TABLE IF NOT EXISTS likes (
            post_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (post_id, user_id)
        );`

	dislikesTable := `
        CREATE TABLE IF NOT EXISTS dislikes (
            post_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (user_id) REFERENCES users(uid),
            PRIMARY KEY (post_id, user_id)
        );`

	sessionTable := `
        CREATE TABLE IF NOT EXISTS sessions (
            session_id INTEGER PRIMARY KEY AUTOINCREMENT,
            session TEXT,
            user_id INTEGER,
            timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(uid)
        );`

	postTable := `
        CREATE TABLE IF NOT EXISTS posts (
            post_id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER,
            dislike INTEGER DEFAULT 0,
            like INTEGER DEFAULT 0,
            post_heading TEXT NOT NULL,
            post_data TEXT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(uid)
        );`

	postCategoriesTable := `
        CREATE TABLE IF NOT EXISTS post_categories (
            post_id INTEGER,
            category_id INTEGER,
            FOREIGN KEY (post_id) REFERENCES posts(post_id),
            FOREIGN KEY (category_id) REFERENCES categories(category_id),
            PRIMARY KEY (post_id, category_id)
        );`

	// Execute table creation statements
	ExecuteSQL(db, userTable)
	ExecuteSQL(db, postTable)
	ExecuteSQL(db, categoryTable)
	ExecuteSQL(db, commentTable)
	ExecuteSQL(db, likeCommentTable)
	ExecuteSQL(db, dislikeCommentTable)
	ExecuteSQL(db, likesTable)
	ExecuteSQL(db, dislikesTable)
	ExecuteSQL(db, sessionTable)
	ExecuteSQL(db, postCategoriesTable)
}

func ExecuteSQL(db *sql.DB, sqlStatement string) {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Printf("SQL Error executing: %s, Error: %v", sqlStatement, err)
		log.Fatal(err)
	}
}

type Post struct {
	ID      int
	Title   string
	Content string
}

func InsertDummyData(db *sql.DB) {
	// Dummy post data
	posts := []Post{
		{Title: "Welcome to the forum", Content: "This is a welcome post!"},
		{Title: "Go programming", Content: "Let's discuss Go programming!"},
		{Title: "Database integration", Content: "Learn how to use SQLite with Go."},
	}

	// Insert each post into the database
	for _, post := range posts {
		_, err := db.Exec("INSERT INTO posts (post_heading, post_data) VALUES (?, ?)", post.Title, post.Content)
		if err != nil {
			log.Printf("Error inserting post: %v, Title: %s", err, post.Title)
			log.Fatal(err)
		}
	}
}
