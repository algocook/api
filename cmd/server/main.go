package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//users "github.com/algocook/proto/users"
	users "algocook/proto/users"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 64)

	client, err := users.NewClient()
	if err != nil {
		json.NewEncoder(w).Encode(users.User{
			Error: err.Error(),
		})
		return
	}

	user := client.GetUser(id)
	json.NewEncoder(w).Encode(user)
}

func getAvailability(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	username := params["username"]

	client, err := users.NewClient()
	if err != nil {
		json.NewEncoder(w).Encode(users.IsAvailable{
			Error: err.Error(),
		})
		return
	}

	result := client.CheckUsername(username)
	json.NewEncoder(w).Encode(result)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var user users.User
	decoder.Decode(&user)
	fmt.Print(user)

	client, err := users.NewClient()
	if err != nil {
		json.NewEncoder(w).Encode(users.User{
			Error: err.Error(),
		})
		return
	}

	user = client.PostUser(user.Username, user.Title, user.Description)
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/usernameavailable/{username}", getAvailability).Methods("GET")
	router.HandleFunc("/user", postUser).Methods("POST")
	http.ListenAndServe(":80", router)
}
