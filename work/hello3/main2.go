package main

var userStore []*user

type user struct {
	ID    string
	Email string
}

func createUser(u *user) user {
	userStore = append(userStore, u)
	return u
}

func listUsers() []user {
	return UserStore
}
