package routes

import (
	"log"

	"github.com/Edw-Castro/Preis/internal/core/services"
	database "github.com/Edw-Castro/Preis/internal/infrastructure/database/mysql"
	"github.com/Edw-Castro/Preis/internal/infrastructure/repositories/sql"
	server "github.com/Edw-Castro/Preis/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func User() func(c *gin.Context) {
	db, err := database.SetupDatabaseUsersConnection()
	if err != nil {
		log.Fatalf("Error cannot connect to db: %v", err)
	}
	repoUser := sql.NewProductRepository(db)
	servUser := services.NewBusinessOwnerService(repoUser)
	getUserEndpoint := server.User(servUser)
	return getUserEndpoint

}
