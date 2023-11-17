package server

import (
	"errors"
	"net/http"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/gin-gonic/gin"
)

/************* Handlers BusinessUsers *************/
func PostBusinessOwnerEndpoint(service ports.BusinessOwnerService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var tp interface{}
		name, password, email := c.Query("name"), c.Query("password"), c.Query("email")
		businessOwner := &domain.User{
			Name:     name,
			Password: password,
			Email:    email,
		}
		err := service.SignUp(businessOwner)
		if err != nil {
			err := errors.New("Ya está registrado")
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, tp)
	}
}
