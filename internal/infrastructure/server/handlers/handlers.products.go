package server

import (
	"net/http"
	"strconv"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
	"github.com/gin-gonic/gin"
)

func GetDetailProductEndpoint(service ports.ProductService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var product domain.Product
		id, isExist := c.Params.Get("id")
		if isExist {
			idNumb, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			product, err = service.Detail(idNumb)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		// fmt.Println("Ya se inscribió el usuario")
		productArray := []domain.Product{product}
		c.JSON(http.StatusOK, dto.BuildResponseProductGet(productArray))
		// c.Redirect(http.StatusFound, "http://localhost:8080/auth/login")
	}
}
