package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go-crud/entity"
	"go-crud/service"
	"io/ioutil"
	"net/http"
)

func SaveUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := bodyUnMarshal(r, &user)
	var result []byte
	w.Header().Set("Content-Type", "application/json")
	if err != nil || user.Id != 0 {
		err = errors.New(fmt.Sprint("Json data is incorrect or user id not empty"))
		result = errorHandler(err, http.StatusBadRequest)
	} else {
		user, er := service.SaveUser(user)
		if er != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorHandler(er, http.StatusBadRequest))
			return
		}
		res, err := unMarshal(user)
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		result = res
		w.WriteHeader(http.StatusCreated)
	}
	_, _ = w.Write(result)
}

func GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	var users []entity.User
	users = service.GetAllUsers()
	result, err := arrayUnMarshal(users)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = w.Write(result)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := service.GetUserById(id)
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := unMarshal(user)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := bodyUnMarshal(r, &user)
	var result []byte
	w.Header().Set("Content-Type", "application/json")
	if err != nil || user.Id == 0 {
		err = errors.New(fmt.Sprint("Json data is incorrect or user id empty or zero"))
		result = errorHandler(err, http.StatusBadRequest)
	} else {
		res, err := service.UpdateUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			result = errorHandler(err, http.StatusBadRequest)
		} else {
			r, e := unMarshal(res)
			if e == nil {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				r = errorHandler(err, http.StatusInternalServerError)
			}
			result = r
		}
	}
	_, _ = w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := service.DeleteUser(id)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorHandler(err, http.StatusBadRequest))
	}
}

func bodyUnMarshal(request *http.Request, user *entity.User) error {
	reqBody, _ := ioutil.ReadAll(request.Body)
	return json.Unmarshal(reqBody, &user)
}

func unMarshal(generic entity.Generic) ([]byte, error) {
	result, err := json.Marshal(generic)
	if err != nil {
		return errorHandler(err, http.StatusInternalServerError), err
	}
	return result, nil
}

func arrayUnMarshal(generic []entity.User) ([]byte, error) {
	result, err := json.Marshal(generic)
	if err != nil {
		return errorHandler(err, http.StatusInternalServerError), err
	}
	return result, nil
}

func errorHandler(err error, status int) []byte {
	issue := entity.Error{Error: err, Cause: err.Error(), Status: status}
	res, _ := json.Marshal(issue)
	return res
}
