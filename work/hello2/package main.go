package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	//create new router
	r := chi.NewRouter()

	//use the middleware

	r.Use(middleware.Logger)

	//Declare your routes
	//r.Get("/", helloHandler)
	r.Get("/api/users", getUsersHandler)
	r.Get("/api/users/read", readUsersHandler)
	r.Get("/api/users/create", createUsersHandler)
	r.Get("/api/users/update", updateUsersHandler)
	r.Get("/api/users/delete", deleteUsersHandler)

	//State that the server is running

	fmt.Println("running on :8080")

	//Run the server
	log.Fatalln(http.ListenAndServe(":8080", r))

	http.ListenAndServe(":8080", r) // Server
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) { // Handler

	fmt.Fprint(w, "Hello World")
}
func readUsersHandler(w http.ResponseWriter, r *http.Request) { // Handler

	fmt.Fprint(w, "Hello World")
}
func createUsersHandler(w http.ResponseWriter, r *http.Request) { // Handler
	fmt.Fprint(w, "Hello World")

}

func updateUsersHandler(w http.ResponseWriter, r *http.Request) { // Handler
	fmt.Fprint(w, "Hello World")

}
func deleteUsersHandler(w http.ResponseWriter, r *http.Request) { // Handler
	fmt.Fprint(w, "Hello World")

}
