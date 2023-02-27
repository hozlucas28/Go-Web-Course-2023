package models

import (
	"GoAPIRESTGORM/db"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

// Migrar objeto a la base de datos
func MigrateUser() {
	db.Database.AutoMigrate(User{})
}
