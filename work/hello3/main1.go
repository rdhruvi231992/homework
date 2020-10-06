package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofrs/uuid"
)

func main() {

	userStore = []*user{}

	// Create a new router
	r := chi.NewRouter()

	// Use middleware logger on each request
	r.Use(middleware.Logger)

	// Declare your routes
	r.Get("/api/users", listUsersHandler)
	r.Post("/api/users/create", createUserHandler)
	r.Post("/api/users/update", updateUserHandler)
	r.Delete("/api/users/delete", deleteUserHandler)

	// State that the server is running
	fmt.Println("Running on :8080")

	//Run the server
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(userStore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// CreateUserRequest ()
type CreateUserRequest struct {
	Email string `json:"email"`
}

// CreateUserResponse ()
type CreateUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// JSON Decode to struct
	req := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := uuid.Must(uuid.NewV4()).String()
	// Business Logic
	newUser := &user{
		ID:    userID,
		Email: req.Email,
	}
	createUser(newUser)

	// JSON Encode from struct
	// Prepare response

	resp := &CreateUserResponse{
		ID:    userID,
		Email: req.Email,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err, Error(), http.StatusBadRequest)
		return
	}
}

func readUserHandler(w http.ResponseWriter, r *http.Request) {

}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
