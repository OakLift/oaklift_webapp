// auth_oauth.go

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

var store = sessions.NewCookieStore([]byte("your-session-key")) // Replace with a secure key

func init() {
	// Set up OAuth providers
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8000/auth/google/callback"),
		github.New(os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_CLIENT_SECRET"), "http://localhost:8000/auth/github/callback"),
		facebook.New(os.Getenv("FACEBOOK_CLIENT_ID"), os.Getenv("FACEBOOK_CLIENT_SECRET"), "http://localhost:8000/auth/facebook/callback"),
	)
}

// Handles the initial OAuth login flow
func oauthLoginHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Store = store
	gothic.BeginAuthHandler(w, r)
}

// Handles the OAuth callback
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Store = store
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE oauth_provider = $1 AND oauth_id = $2", user.Provider, user.UserID).Scan(&userID)
	if err == sql.ErrNoRows {
		// User does not exist, create a new one
		_, err = db.Exec("INSERT INTO users (oauth_provider, oauth_id, email, first_name, last_name, role) VALUES ($1, $2, $3, $4, $5, $6)",
			user.Provider, user.UserID, user.Email, user.FirstName, user.LastName, "mentee") // Default role, can be changed
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}
		err := db.QueryRow("SELECT id FROM users WHERE oauth_provider = $1 AND oauth_id = $2", user.Provider, user.UserID).Scan(&userID)
		if err != nil {
			http.Error(w, "Error getting row", http.StatusInternalServerError)
			return
		}
	}

	session, _ := store.Get(r, "auth-session")
	session.Values["user_id"] = userID
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User logged in successfully with %s\n", user.Provider)
}

// Handles OAuth logout
func oauthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Store = store
	gothic.Logout(w, r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User logged out successfully")
}
