/* --------------------------------------------------------------------------
 * APUNTES:
 * 		   Un "mux" es una ruta asociada a un "handler".
 *
 *
 * IMPORTANTE:
 *  			  - <go run github.com/pilu/fresh> = Levanta el servidor en
 *													 modo desarrollo, actualizándose
 *													 automaticamente por cada cambio
 *													 que se genere.
-------------------------------------------------------------------------- */

package main

import (
	"fmt"
	"log"
	"net/http"
)

/* -------------------------------- Handlers -------------------------------- */

func Hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El método es: " + r.Method)
	fmt.Fprintln(rw, "¡Hola Mundo!")
}

func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La página no funciona: Status 404", http.StatusNotFound)
}

func Greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)          // Obtengo la ruta
	fmt.Println(r.URL.RawQuery) // Obtengo las consultas
	fmt.Println(r.URL.Query())  // Obtengo las consultas en un mapa

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(rw, "• Nombre: %s\n• Edad: %s\n", name, age)
}

/* -------------------------------- Servidor -------------------------------- */

func main() {
	// Muxs
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)                    // http://localhost:3000/
	mux.HandleFunc("/pagenotfound", PageNotFound) // http://localhost:3000/pagenotfound
	mux.HandleFunc("/error", Error)               // http://localhost:3000/error
	mux.HandleFunc("/greet", Greet)               // http://localhost:3000/greet?name=Lucas&age=20

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor esta funcionando en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
