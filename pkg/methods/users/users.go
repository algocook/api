package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/algocook/proto/users"
	"github.com/gorilla/mux"
)

// GetOne возвращает 1 пользователя
func GetOne(w http.ResponseWriter, r *http.Request) {
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

// GetUsernameAvailability проверяет доступен ли никнейм
func GetUsernameAvailability(w http.ResponseWriter, r *http.Request) {
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

// PostOne загружает 1 пользователя в бд
func PostOne(w http.ResponseWriter, r *http.Request) {
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

// DeleteOne удаляет из бд
func DeleteOne(w http.ResponseWriter, r *http.Request) {
}

// Search поиск по пользователям
func Search(w http.ResponseWriter, r *http.Request) {
}
