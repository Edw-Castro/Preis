package database

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // En tu package http.go
// /*func SetupDatabaseProductsConnection() (*gorm.DB, error) {
// 	// Configura la cadena de conexión a la base de datos
// 	//db, err := gorm.Open("mysql", "usuario:contraseña@tcp(direccion_base_de_datos)/nombre_base_de_datos?parseTime=true")
// 	dsn := "host=localhost user=postgres dbname=products sslmode=disable password=postgres"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }*/
// func SetupDatabaseClientConnection() (*gorm.DB, error) {
// 	dsn := "root:password@tcp(ip:3306)/database1?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Error al conectar a la base de datos: " + err.Error())
// 	}

// 	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	fmt.Println("SetupDatabaseClientConnection", db)
// 	return db, nil
// }
