package routes

import (
	"log"

	"github.com/Edw-Castro/Preis/internal/core/services"
	database "github.com/Edw-Castro/Preis/internal/infrastructure/database/mysql"
	"github.com/Edw-Castro/Preis/internal/infrastructure/repositories/sql"
	server "github.com/Edw-Castro/Preis/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func ProductDetail() func(c *gin.Context) {
	db, err := database.SetupDatabaseArticleConnection()
	if err != nil {
		log.Fatalf("Error cannot connect to db: %v", err)
	}
	repoProduct := sql.NewProductRepository(db)
	servProduct := services.NewProductService(repoProduct)
	getProductEndpoint := server.GetDetailProductEndpoint(servProduct)
	return getProductEndpoint
}
