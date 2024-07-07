package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// User struct to map user data
type User struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password,omitempty"`
	Email         string    `json:"email"`
	OAuthProvider string    `json:"oauth_provider,omitempty"`
	OAuthID       string    `json:"oauth_id,omitempty"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Role          string    `json:"role"`
	Bio           string    `json:"bio,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Connect to the database
func initDB() {
	var err error
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Connected to database")
	}
}

// Register a new user
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Unable to hash password")
			http.Error(w, "Server error, unable to create your account.", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
	}

	log.Print(user)
	sqlStatement := `INSERT INTO users (username, password, email, oauth_provider, oauth_id, first_name, last_name, role, bio)
                     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(sqlStatement, user.Username, user.Password, user.Email, user.OAuthProvider, user.OAuthID, user.FirstName, user.LastName, user.Role, user.Bio)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Server error, unable to create your account.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login a user
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var dbUser User

	json.NewDecoder(r.Body).Decode(&user)

	sqlStatement := `SELECT id, username, password, email, oauth_provider, oauth_id, first_name, last_name, role, bio, created_at, updated_at
                     FROM users WHERE username=$1`
	row := db.QueryRow(sqlStatement, user.Username)
	err := row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password, &dbUser.Email, &dbUser.OAuthProvider, &dbUser.OAuthID, &dbUser.FirstName, &dbUser.LastName, &dbUser.Role, &dbUser.Bio, &dbUser.CreatedAt, &dbUser.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username", http.StatusUnauthorized)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	if user.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dbUser)
}

func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/auth/{provider}/callback", oauthCallbackHandler)
	router.HandleFunc("/auth/{provider}", oauthLoginHandler)
	router.HandleFunc("/logout", oauthLogoutHandler)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
