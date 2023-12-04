package server

import (
	"net/http"

	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
	"github.com/gin-gonic/gin"
)

func Compare(service ports.ArticlePriceService) func(*gin.Context) {
	return func(c *gin.Context) {
		products, err := service.Compare(c.Query("name"), c.Query("brand"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		c.JSON(http.StatusOK, dto.BuildResponseArticlePrice(products))
	}
}
