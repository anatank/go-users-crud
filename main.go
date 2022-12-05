package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID      string `json: "id"`
	Name    string `json: "name"`
	Email   string `json: "email"`
	Country string `json: "country"`
	Phone   string `json: "phone"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode((item))
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()

	users = append(users, User{
		ID:      "1",
		Name:    "Sybill Cruz",
		Email:   "ullamcorper.eu@icloud.net",
		Country: "Ukraine",
		Phone:   "(372) 958-6641",
	})
	users = append(users, User{
		ID:      "2",
		Name:    "Sybill Cruz",
		Email:   "turpis@protonmail.ca",
		Country: "India",
		Phone:   "(566) 745-7729",
	})
	users = append(users, User{
		ID:      "3",
		Name:    "Lunea David",
		Email:   "pretium.aliquet@icloud.edu",
		Country: "France",
		Phone:   "(799) 733-5820",
	})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Print("Starting at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
