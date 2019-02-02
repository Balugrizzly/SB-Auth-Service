package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Env struct {
	db *gorm.DB
}

func main() {

	// db Init
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// db migrations
	autoMigration(db)

	env := &Env{db: db}

	// Routes
	r := mux.NewRouter()
	r.HandleFunc("/createuser", env.createUser).Methods("POST")

	// Server
	log.Fatal(http.ListenAndServe(":5000", r))

}

func (env *Env) createUser(w http.ResponseWriter, r *http.Request) {
	// creates a new user in the db
	// returns the created user on success

	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// check if username already exists
	usernameCount := 0
	env.db.Model(&User{}).Where("name = ?", user.Name).Count(&usernameCount)

	if usernameCount > 0 {
		_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "Name already taken"})
		return
	}

	// check if username is at least one char
	if len(user.Name) == 0 {
		_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "No name provided"})
		return
	}

	// hash pw
	user.Pw, _ = hashPassword(user.Pw)

	env.db.Create(&user)
	// return user
	_ = json.NewEncoder(w).Encode(&user)

}
