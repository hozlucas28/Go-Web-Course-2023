package models

import (
	"GoMySQL/db"
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
	Id       int64
	UserName string
	Password string
	Email    string
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
func GetUser(id int) *User {
	user := NewUser("", "", "")
	query := "SELECT id, userName, password, email FROM users WHERE id=?"
	rows, _ := db.Query(query, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}
	return user
}

// Listar
func ListUsers() Users {
	users := Users{}
	query := "SELECT id, userName, password, email FROM users"
	rows, _ := db.Query(query)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}
