package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// fack db
var users []User

// middlewares, helper file
func isEmpty(u *User) bool {
	return u.FullName == "" && u.Email == "" && u.Password == ""
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<font color='#00ffab'><h1>Welcome to rest API</h1></font>"))
}

func serveUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func serveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found!")
	return
}

func serveAddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	var i int = 1
	json.NewDecoder(r.Body).Decode(&user)
	if r.Body != nil && !isEmpty(&user) {
		user.ID = strconv.Itoa(i)
		users = append(users, user)
		json.NewEncoder(w).Encode(user)
		i++
		return
	}

	json.NewEncoder(w).Encode("user data not found!")
}

func serverUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:], users[index+1:]...)
			var user User
			json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}

		json.NewEncoder(w).Encode("User not found to be updated!")
	}
}

func serveDeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("User has been deleted!")

}

func main() {
	fmt.Println("Running od 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/users", serveUsers)
	r.HandleFunc("/user/{id}", serveUser)
	r.HandleFunc("/add", serveAddUser)
	r.HandleFunc("/update/{id}", serverUpdateUser)
	r.HandleFunc("/delete/{id}", serveDeleteUser)
	http.ListenAndServe(":8080", r)
}

func fetal(err error) {
	if err != nil {
		panic(err)
	}
}
