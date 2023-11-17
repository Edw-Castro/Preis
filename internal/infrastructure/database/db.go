package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// En tu package http.go
/*func SetupDatabaseProductsConnection() (*gorm.DB, error) {
	// Configura la cadena de conexión a la base de datos
	//db, err := gorm.Open("mysql", "usuario:contraseña@tcp(direccion_base_de_datos)/nombre_base_de_datos?parseTime=true")
	dsn := "host=localhost user=postgres dbname=products sslmode=disable password=postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}*/
func SetupDatabaseClientConnection() (*gorm.DB, error) {
	// Configura la cadena de conexión a la base de datos
	//db, err := gorm.Open("mysql", "usuario:contraseña@tcp(direccion_base_de_datos)/nombre_base_de_datos?parseTime=true")
	dsn := "host=localhost user=postgres dbname=users sslmode=disable password=postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
