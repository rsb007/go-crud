package main

import (
	"awesomeProject/entity"
	"awesomeProject/utils"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	db, _ := utils.DbConnection()
	db.AutoMigrate(&entity.User{})
	myRouter := mux.NewRouter().StrictSlash(true)
	port := os.Getenv("port")
	if port == "" {
		port = "3000"
		log.Printf("Defaulting to port %s", port)
	}
	Routes(myRouter)
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}
