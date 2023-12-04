package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseUsersConnection() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3306)/database1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	fmt.Println("SetupDatabaseClientConnection", db)
	fmt.Println("CONECTADA LA BASE DE DATOS 2")
	return db, nil
}
