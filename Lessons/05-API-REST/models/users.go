package models

import (
	"GoAPIREST/db"
	"fmt"
)

/* --------------------------- Variables Globales --------------------------- */

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	userName VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(30),
	dateCreated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

/* --------------------------------- Objetos -------------------------------- */

type User struct {
	// JSON
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	// XML
	/*
		Id       int64  `xml:"id"`
		UserName string `xml:"username"`
		Password string `xml:"password"`
		Email    string `xml:"email"`
	*/

	// YAML
	/*
		Id       int64  `yaml:"id"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
	*/
}

type Users []User

// Constructor - User
func NewUser(name, password, email string) *User {
	user := &User{
		UserName: name,
		Password: password,
		Email:    email,
	}

	return user
}

/* --------------------------------- Métodos -------------------------------- */

// Guardar o Editar
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Insertar
func (user *User) insert() {
	query := "INSERT users SET userName=?, password=?, email=?"
	result, _ := db.Exec(query, user.UserName, user.Password, user.Email)

	user.Id, _ = result.LastInsertId()
}

// Eliminar
func (user *User) Delete() {
	query := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(query, user.Id)

	if err != nil {
		fmt.Println(err)
	}
}

// Actualizar
func (user *User) update() {
	query := "UPDATE users SET userName=?, password=?, email=? WHERE id=?"
	_, err := db.Exec(query, user.UserName, user.Password, user.Email, user.Id)

	if err != nil {
		fmt.Println(err)
	}
}

/* -------------------------------- Funciones ------------------------------- */

// Crear e insertar
func CreateUser(name, password, email string) *User {
	user := NewUser(name, password, email)
	user.Save()
	return user
}

// Obtener un usuario
func GetUser(id int) (*User, error) {
	user := NewUser("", "", "")
	query := "SELECT id, userName, password, email FROM users WHERE id=?"

	if rows, err := db.Query(query, id); err == nil {
		for rows.Next() {
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		}
		return user, nil
	} else {
		return nil, err
	}
}

// Listar
func ListUsers() (Users, error) {
	users := Users{}
	query := "SELECT id, userName, password, email FROM users"

	if rows, err := db.Query(query); err == nil {
		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
			users = append(users, user)
		}
		return users, nil
	} else {
		return nil, err
	}
}
