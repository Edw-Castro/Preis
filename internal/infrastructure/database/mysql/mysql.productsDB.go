package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseArticleConnection() (*gorm.DB, error) {
	// Configura la cadena de conexión a la base de datos
	//db, err := gorm.Open("mysql", "usuario:contraseña@tcp(direccion_base_de_datos)/nombre_base_de_datos?parseTime=true")
	dsn := "root:password@tcp(localhost:3306)/database2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("CONECTADA LA BASE DE DATOS 1")
	return db, nil
}
