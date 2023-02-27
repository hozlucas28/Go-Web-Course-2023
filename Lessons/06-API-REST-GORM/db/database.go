package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realizar conexión a la base de datos
var dsn = "root:289137tx@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		fmt.Println("Conexión exitosa")
		return db
	} else {
		fmt.Println("¡Error en la conexión!", err)
		panic(err)
	}
}()
