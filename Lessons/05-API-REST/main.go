package main

import (
	"GoAPIREST/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Rutas
	mux := mux.NewRouter()

	// EndPoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")                 // Devolver todos los usuarios.
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")       // Devolver usuario especifico.
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")              // Registrar usuario.
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")    // Eliminar usuario.
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE") // Editar usuario.

	// Inicializar servidor
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
