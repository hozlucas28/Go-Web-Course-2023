package handlers

import (
	"GoAPIREST/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Devolver todos los usuarios
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	/* ---------------------------------- Viejo --------------------------------- */

	/*
		db.Connect()
		users, _ := models.ListUsers()
		db.Close()
	*/

	// JSON
	/*
		rw.Header().Set("content-type", "application/json") // Header

		output, _ := json.Marshal(users) // Transformar objeto a tipo JSON.
		outputStr := string(output)      // Casteo de []byte a []string.
		fmt.Fprintln(rw, outputStr)
	*/

	// XML
	/*
		rw.Header().Set("content-type", "text/xml") // Header

		output, _ := xml.Marshal(users) // Transformar objeto a tipo XML.
		outputStr := string(output)     // Casteo de []byte a []string.
		fmt.Fprintln(rw, outputStr)
	*/

	// YALM
	/*
		output, _ := yaml.Marshal(users) // Transformar objeto a tipo YALM.
		outputStr := string(output)      // Casteo de []byte a []string.
		fmt.Fprintln(rw, outputStr)
	*/

	/* ------------------------------ Refactorizado ----------------------------- */

	if users, err := models.ListUsers(); err == nil {
		models.SendData(rw, users)
	} else {
		models.SendNotFound(rw)
	}
}

// Devolver usuario especifico
func GetUser(rw http.ResponseWriter, r *http.Request) {
	/* ---------------------------------- Viejo --------------------------------- */

	/*
		// Obtengo las variables
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"]) // Obtengo y casteo el Id.

		db.Connect()
		user, _ := models.GetUser(id)
		db.Close()

		rw.Header().Set("content-type", "application/json")
		output, _ := json.Marshal(user)
		outputStr := string(output)
		fmt.Fprintln(rw, outputStr)
	*/

	/* ------------------------------ Refactorizado ----------------------------- */

	if user, err := getUserByRequest(r); err == nil {
		models.SendData(rw, user)
	} else {
		models.SendNotFound(rw)
	}
}

// Registrar usuario
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	/* ---------------------------------- Viejo --------------------------------- */

	/*
		// Construyo el usuario
		user := models.User{}
		decoder := json.NewDecoder(r.Body) // Transformar respuesta JSON a objeto.
		if err := decoder.Decode(&user); err == nil {
			db.Connect()
			user.Save()
			db.Close()
		} else {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		}

		rw.Header().Set("content-type", "application/json")
		output, _ := json.Marshal(user)
		outputStr := string(output)
		fmt.Fprintln(rw, outputStr)
	*/

	/* ------------------------------ Refactorizado ----------------------------- */

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err == nil {
		user.Save()
		models.SendData(rw, user)
	} else {
		models.SendUnprocessableEntity(rw)
	}
}

// Eliminar usuario
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	/* ---------------------------------- Viejo --------------------------------- */

	/*
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		db.Connect()
		user, _ := models.GetUser(id)
		user.Delete()
		db.Close()

		rw.Header().Set("content-type", "application/json")
		output, _ := json.Marshal(user)
		outputStr := string(output)
		fmt.Fprintln(rw, outputStr)
	*/

	/* ------------------------------ Refactorizado ----------------------------- */

	if user, err := getUserByRequest(r); err == nil {
		user.Delete()
		models.SendData(rw, user)
	} else {
		models.SendNotFound(rw)
	}
}

// Editar usuario
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	/* ---------------------------------- Viejo --------------------------------- */

	/*
		user := models.User{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err == nil {
			db.Connect()
			user.Save()
			db.Close()
		} else {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		}

		rw.Header().Set("content-type", "application/json")
		output, _ := json.Marshal(user)
		outputStr := string(output)
		fmt.Fprintln(rw, outputStr)
	*/

	/* ------------------------------ Refactorizado ----------------------------- */

	var userId int64

	if user, err := getUserByRequest(r); err == nil {
		userId = user.Id
	} else {
		models.SendNotFound(rw)
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err == nil {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	} else {
		models.SendNotFound(rw)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(id); err == nil {
		return *user, nil
	} else {
		return *user, err
	}
}
