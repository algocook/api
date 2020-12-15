package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/algocook/proto"
	"github.com/gorilla/mux"
)

// UsersClient comment
type UsersClient struct {
	conn *grpc.ClientConn
}

var usersClient UsersClient

// User struct
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Title    string `json:"title"`
}

// GetUser function
func (client *UsersClient) GetUser(id int64) User {
	cli := pb.NewUsersClient(client.conn)
	var request pb.GetUserRequest
	request.Id = id
	response, err := cli.GetUser(context.Background(), &request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return User{
		ID:       response.Id,
		Username: response.Username,
		Title:    response.Title,
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 0, 64)

	user := usersClient.GetUser(id)

	json.NewEncoder(w).Encode(user)
}

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	//args := os.Args
	con, err := grpc.Dial("users:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	usersClient.conn = con

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	http.ListenAndServe(":80", router)
}
