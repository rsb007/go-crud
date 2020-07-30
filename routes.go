package main

import (
	"github.com/gorilla/mux"
	"go-crud/api"
)

const (
	user string = "/user"
)

func Routes(myRouter *mux.Router) {
	//myRouter.HandleFunc("/", indexHandler)
	myRouter.HandleFunc(user, api.SaveUser).Methods("POST")
	myRouter.HandleFunc(user, api.GetAllUsers).Methods("GET")
	myRouter.HandleFunc(user+"/{id}", api.GetUserById).Methods("GET")
	myRouter.HandleFunc(user+"/{id}", api.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc(user, api.UpdateUser).Methods("put")
}
