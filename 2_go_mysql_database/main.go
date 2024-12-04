package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Application struct {
	DB *sql.DB
}

func (app *Application) connectDB() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_mysql_database?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	app.DB = db
	return db, nil
}

func (app *Application) CreateTable(w http.ResponseWriter, r *http.Request) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	_, err := app.DB.Exec(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Create table error: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Users table created successfully!")
}

func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	username := "Hudayberdi"
	password := "secret"
	createdAt := time.Now()

	result, err := app.DB.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Insert user error: %v", err), http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve last insert ID: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Last inserted user's ID: %d", userID)
}

func (app *Application) getUserByID(w http.ResponseWriter, r *http.Request) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := `SELECT id, username, password, created_at FROM users WHERE id=?`
	err = app.DB.QueryRow(query, userID).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("User not found for ID: %d", userID), http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		}
		return
	}

	data := fmt.Sprintf("User Data:\nID: %d\nUsername: %s\nPassword: %s\nCreated At: %v",
		id, username, password, createdAt.Format("2006-01-02 15:04:05"))

	fmt.Fprintln(w, data)
}

func (app Application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, username, password, created_at FROM users`

	type User struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
	}

	var users []User
	rows, err := app.DB.Query(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}

		// log.Printf("User fetched: ID=%d, Username=%s, Password=%s, CreatedAt=%v", user.ID, user.Username, user.Password, user.CreatedAt)

		users = append(users, user)
	}
	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func main() {
	var app Application

	database, err := app.connectDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer database.Close()

	router := mux.NewRouter()
	router.HandleFunc("/create-table", app.CreateTable)
	router.HandleFunc("/users/create", app.CreateUser)
	router.HandleFunc("/users/{id}", app.getUserByID)
	router.HandleFunc("/users", app.getAllUsers)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
