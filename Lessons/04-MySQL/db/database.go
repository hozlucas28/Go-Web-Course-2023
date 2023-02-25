package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:289137tx@tcp(localhost:3306)/goweb_db" // URL a la base de datos
var db *sql.DB                                           // Conexión de la base de datos

// Inicializar conexión
func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	db = connection
	fmt.Println("Conexión establecida.")
}

// Cerrar conexión
func Close() {
	db.Close()
	fmt.Println("Conexión cerrada.")
}

// Revisar conexión
func CheckConnection() bool {
	err := db.Ping()

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// Crear tabla
func CreateTable(schema string, tName string) {
	if !CheckTable(tName) {
		_, err := Exec(schema)

		if err != nil {
			fmt.Println(err)
		}
	}
}

// Revisar la tabla
func CheckTable(tName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tName)
	rows, err := Query(query)

	if err != nil {
		fmt.Println(err)
	}

	return rows.Next()
}

// Vaciar registros de la tabla
func TruncateTable(tName string) {
	query := fmt.Sprintf("TRUNCATE %s", tName)
	_, err := Exec(query)

	if err != nil {
		fmt.Println(err)
	}
}

// Polimorfismo del Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo del Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
