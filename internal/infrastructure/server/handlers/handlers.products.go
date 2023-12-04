package server

import (
	"fmt"
	"net/http"

	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
	"github.com/gin-gonic/gin"
)

func GetDetailProductEndpoint(service ports.ProductService) func(c *gin.Context) {
	return func(c *gin.Context) {
		// id, isExist := c.Params.Get("id")
		product, err := service.GetAllProducts()
		// if isExist {
		// 	idNumb, err := strconv.Atoi(id)
		// 	if err != nil {
		// 		c.AbortWithStatus(http.StatusNotFound)
		// 		return
		// 	}
		// 	if err != nil {
		// 		c.AbortWithStatus(http.StatusNoContent)
		// 		return
		// 	}
		// }
		// fmt.Println("Ya se inscribió el usuario")
		if err != nil {
			fmt.Println("Error en GetDetailProductEndpoint")
		}
		// productArray := []domain.Product{product}
		c.JSON(http.StatusOK, dto.BuildResponseProductGet(product))
		// c.Redirect(http.StatusFound, "http://localhost:8080/auth/login")
	}
}
