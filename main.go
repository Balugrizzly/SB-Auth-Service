package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Env struct {
	db *gorm.DB
}

var (
	userSessions []UserSession
)

func main() {

	// db Init
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.LogMode(true)

	// db migrations
	autoMigration(db)

	env := &Env{db: db}

	// Routes
	r := mux.NewRouter()
	// Post
	r.HandleFunc("/createuser", env.createUser).Methods("POST")
	r.HandleFunc("/authuser", env.authUser).Methods("POST")
	r.HandleFunc("/updateuser", env.updateUser).Methods("POST")
	r.HandleFunc("/deleteuser", env.deleteUser).Methods("POST")

	// All request methods
	r.HandleFunc("/getuser", env.getUser)
	// r.HandleFunc("/usernameexists", env.getUser)

	// Server
	log.Fatal(http.ListenAndServe(":5000", r))

}

func (env *Env) updateUser(w http.ResponseWriter, r *http.Request) {
	// updates the user
	// returns the updated user on success
	w.Header().Set("Content-Type", "application/json")

	// get reqtoken
	reqToken := r.Header.Get("token")

	// check if token is valid
	for _, session := range userSessions {
		if session.SessionToken == reqToken {

			var reqUser User
			_ = json.NewDecoder(r.Body).Decode(&reqUser)

			var dBUser User
			env.db.Where("id = ?", session.User.ID).Find(&dBUser)

			if dBUser.Name != reqUser.Name {
				if usernameExists(env.db, reqUser.Name) {
					_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "Username already exists"})
					return
				}
			}

			if reqUser.Name != "" {
				dBUser.Name = reqUser.Name
			}
			if reqUser.Pw != "" {
				dBUser.Pw, _ = hashPassword(reqUser.Pw)
			}
			env.db.Save(&dBUser)

			_ = json.NewEncoder(w).Encode(&dBUser)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "Auth failed"})
}

func (env *Env) deleteUser(w http.ResponseWriter, r *http.Request) {
	// soft deletes the user
	// returns status true on success
	w.Header().Set("Content-Type", "application/json")
	// get token
	reqToken := r.Header.Get("token")

	// check if token matches
	for i, session := range userSessions {
		if session.SessionToken == reqToken {
			fmt.Fprintf(w, "{%q: %v}", "status", true)

			// delete user
			var dBUser User
			env.db.Where("id = ?", session.User.ID).Find(&dBUser)
			env.db.Delete(&dBUser)

			userSessions = removeSession(userSessions, i)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "Auth failed"})
}

func (env *Env) getUser(w http.ResponseWriter, r *http.Request) {

}

func (env *Env) authUser(w http.ResponseWriter, r *http.Request) {
	// authenticates the user
	// returns a jwt on success
	w.Header().Set("Content-Type", "application/json")

	var reqUser User
	_ = json.NewDecoder(r.Body).Decode(&reqUser)

	var dBUser User
	env.db.Where("name = ?", reqUser.Name).Or("id = ?", reqUser.ID).Find(&dBUser)

	if !checkPasswordHash(reqUser.Pw, dBUser.Pw) {
		_ = json.NewEncoder(w).Encode(ErrorResponse{Msg: "Auth failed"})
		return
	}

	token := generateToken()
	// return token
	fmt.Fprintf(w, "{%q: %q}", "token", token)

	// update user sessions
	sessionExist := false
	for _, session := range userSessions {
		if session.User.ID == dBUser.ID {
			sessionExist = true
			session.User = dBUser
			session.SessionToken = token
			session.LoginTimeUnix = time.Now().Unix()
			session.LastSeenUnix = time.Now().Unix()
		}
	}

	if !sessionExist {
		userSession := UserSession{}
		userSession.User = dBUser
		userSession.SessionToken = token
		userSession.LoginTimeUnix = time.Now().Unix()
		userSession.LastSeenUnix = time.Now().Unix()
		userSessions = append(userSessions, userSession)
	}
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
