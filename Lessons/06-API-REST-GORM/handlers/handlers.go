package handlers

import (
	"GoAPIRESTGORM/db"
	"GoAPIRESTGORM/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Devolver todos los usuarios
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users) // Recuperar usuarios (registros).
	sendData(rw, users, http.StatusOK)
}

// Devolver usuario especifico
func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err == nil {
		sendData(rw, user, http.StatusOK)
	} else {
		sendError(rw, http.StatusNotFound)
	}
}

// Registrar usuario
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err == nil {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	} else {
		sendError(rw, http.StatusUnprocessableEntity)
	}
}

// Eliminar usuario
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err == nil {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	} else {
		sendError(rw, http.StatusNotFound)
	}
}

// Editar usuario
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64

	if oldUser, err := getUserById(r); err == nil {
		userId = oldUser.Id

		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err == nil {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		} else {
			sendError(rw, http.StatusUnprocessableEntity)
		}
	} else {
		sendError(rw, http.StatusNotFound)
	}
}

// Obtener usuario especifico
func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	if err := db.Database.First(&user, userId); err.Error == nil {
		return user, nil
	} else {
		return user, err
	}
}
