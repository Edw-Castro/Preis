package routes

import (
	"log"

	"github.com/Edw-Castro/Preis/internal/core/services"
	database "github.com/Edw-Castro/Preis/internal/infrastructure/database/mysql"
	"github.com/Edw-Castro/Preis/internal/infrastructure/repositories/sql"
	server "github.com/Edw-Castro/Preis/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func ProductCompare() func(*gin.Context) {
	db1, err := database.SetupDatabaseArticleConnection()
	db2, err := database.SetupDatabaseUsersConnection()
	if err != nil {
		log.Fatalf("Error cannot connect to db: %v", err)
	}
	repoArticlePrice := sql.NewArticleRepository(db1, db2)
	servProduct := services.NewArticlePriceService(repoArticlePrice)
	getProductEndpoint := server.Compare(servProduct)
	return getProductEndpoint
}
