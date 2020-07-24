package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB
var err error

type User struct {
	Id, Age                                          int32
	FirstName, LastName, Email, City, State, Country string
	CreatedAt, ModifiedAt                            time.Time
}

func dbConnection() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:9090)/golang")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
}

func saveUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "json data is incorrect", http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()
	db.Create(&user)
	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(result)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var user User
	db.Model(user).Where("id = ?", key).Find(&user)
	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)

}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	var users []User
	db.Find(&users)
	result, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "json data is incorrect", http.StatusInternalServerError)
		return
	}
	var persistedUser User
	db.Model(user).Where("id = ?", user.Id).Find(&persistedUser)
	if persistedUser.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		message := fmt.Sprintf("User with id = %d not exist", user.Id)
		_, _ = w.Write([]byte(message))
		return
	}
	user.ModifiedAt = time.Now()
	db.Model(user).Where("id = ?", user.Id).Update(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db.Where("id = ?", id).Delete(&User{})
	w.WriteHeader(http.StatusNoContent)
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func handelRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", index)
	myRouter.HandleFunc("/user", saveUser).Methods("POST")
	myRouter.HandleFunc("/user", getAllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{id}", getUserById).Methods("GET")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user", updateUser).Methods("put")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}
func main() {
	dbConnection()
	db.AutoMigrate(&User{})
	handelRequest()
}
