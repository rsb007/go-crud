package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go-crud/entity"
	"go-crud/utils"
	"time"
)

func SaveUser(user entity.User) (entity.User, error) {
	db, _ := utils.DbConnection()
	user.CreatedAt = time.Now()
	user.ModifiedAt = time.Now()
	checkUser := findUserByEmail(user.Email, db)
	if checkUser.Id != 0 {
		return user, errors.New(fmt.Sprintf("User with email %s already exist", user.Email))
	}
	db.Create(&user)
	defer utils.DbCloseConnection(db)
	return user, nil
}

func findUserByEmail(email string, db *gorm.DB) entity.User {
	var checkUser entity.User
	db.Model(email).Where("email = ?", email).Find(&checkUser)
	return checkUser
}

func GetUserById(id string) entity.User {
	db, _ := utils.DbConnection()
	var user entity.User
	db.Model(user).Where("id = ?", id).Find(&user)
	defer utils.DbCloseConnection(db)
	return user
}

func GetAllUsers() []entity.User {
	db, _ := utils.DbConnection()
	var users []entity.User
	db.Find(&users)
	defer utils.DbCloseConnection(db)
	return users
}

func UpdateUser(user *entity.User) (entity.User, error) {
	db, _ := utils.DbConnection()
	checkUser := findUserByEmail(user.Email, db)
	if checkUser.Id == 0 {
		return *user, errors.New(fmt.Sprintf("user with email=%s not exits", user.Email))
	}
	user.ModifiedAt = time.Now()
	db.Model(&checkUser).Where("email = ? and id = ?", user.Email, user.Id).Update(user)
	db.Save(&user)
	defer utils.DbCloseConnection(db)
	return *user, nil
}

func DeleteUser(id string) error {
	db, _ := utils.DbConnection()
	checkUser := GetUserById(id)
	if checkUser.Id == 0 {
		return errors.New(fmt.Sprintf("user with id =%s not exits", id))
	}
	db.Where("id = ?", id).Delete(&entity.User{})
	defer utils.DbCloseConnection(db)
	return nil
}
