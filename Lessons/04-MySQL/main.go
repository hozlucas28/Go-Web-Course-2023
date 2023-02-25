package main

import (
	"GoMySQL/db"
)

func main() {
	// Iniciar conexión
	db.Connect()
	db.CheckConnection()

	// Crear tabla
	/*
		db.CreateTable(models.UserSchema, "users")
		fmt.Println(db.CheckTable("users"))
	*/

	// Crear e Insertar usuario
	/*
		user := models.CreateUser("Nahuel", "112233", "nahuel@gmail.com")
		fmt.Println(user)
	*/

	// Mostrar usuarios
	/*
		users := models.ListUsers()
		fmt.Println(users)
	*/

	// Obtener usuario por ID
	/*
		user2 := models.GetUser(2)
		fmt.Println(user2)
	*/

	// Modificar usuario
	/*
		user2 := models.GetUser(2)
		user2.UserName = "Matias"
		user2.Save()

		users = models.ListUsers()
		fmt.Println(users)
	*/

	// Borrar usuario
	/*
		user2 := models.GetUser(2)
		specificUser.Delete()

		users = models.ListUsers()
		fmt.Println(users)
	*/

	// Vaciar tabla
	/*
		db.TruncateTable("users")

		users = models.ListUsers()
		fmt.Println(users)
	*/

	// Cerrar conexión
	db.Close()
	db.CheckConnection()
}
