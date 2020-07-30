package utils

import (
	"github.com/jinzhu/gorm"
	"log"
)

func DbConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:9090)/golang?parseTime=true")
	if err != nil {
		log.Println("Connection Failed to Open")
		return db, err
	} else {
		log.Println("Connection Established")
		return db, nil
	}
}

func DbCloseConnection(db *gorm.DB) {
	_ = db.Close()
}
