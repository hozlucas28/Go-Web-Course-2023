package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Active  bool
	IsAdmin bool
	Courses []Course
}

type Course struct {
	Name string
}

func Greet(name string) string {
	return "¡Hola " + name + " desde una función!"
}

// Herencia, variables, datos, operadores e iteración en template
func Index(rw http.ResponseWriter, r *http.Request) {
	template, error := template.ParseFiles("templates/index.html", "templates/base.html")

	c1 := Course{Name: "Curso de HTML"}
	c2 := Course{Name: "Curso de CSS"}
	c3 := Course{Name: "Curso de GO"}
	c4 := Course{Name: "Curso de Java"}
	courses := []Course{c1, c2, c3, c4}

	user1 := User{Name: "Lucas", Age: 20, Active: true, IsAdmin: false, Courses: courses}

	if error == nil {
		template.Execute(rw, user1) // Llamado al template.
	} else {
		panic(error)
	}
}

// Herencia y función en template
func Function(rw http.ResponseWriter, r *http.Request) {
	functions := template.FuncMap{"greet": Greet}
	// Primer método - Dependiente al manejo de errores
	/*
		template, error := template.New("function.html").Funcs(functions).ParseFiles("templates/function.html", "templates/base.html")

		if error == nil {
			template.Execute(rw, nil)
		} else {
			panic(error)
		}
	*/

	// Segundo método - Independiente al manejo de errores
	template := template.Must(template.New("function.html").Funcs(functions).ParseFiles("templates/function.html", "templates/base.html"))
	template.Execute(rw, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)            // http://localhost:3000/
	mux.HandleFunc("/function", Function) // http://localhost:3000/function

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor esta funcionando en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
