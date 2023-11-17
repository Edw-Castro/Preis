package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/infrastructure/server/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Clave secreta para firmar tokens JWT (mejorarlo para obtnerlo desde env)
var secretKey = []byte("tu_clave_secreta")

func Register(service ports.BusinessOwnerService) func(*gin.Context) {
	return func(c *gin.Context) {

		var registerUser dto.ResponseUserGet

		err := c.ShouldBind(&registerUser)

		if err != nil {
			log.Fatalf("El error en el register es %v", err)
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("nombre", registerUser.Name, "apellido", registerUser.LastName, "email", registerUser.Email, "password: ", registerUser.Password)
		fmt.Println(registerUser.Email)
		businessUser := domain.User{
			Email:    registerUser.Email,
			Password: registerUser.Password,
			// Asigna otros valores según sea necesario
			Name: registerUser.Name,

			LastName: registerUser.LastName,
			Role:     "Rol por defecto",
		}
		fmt.Println("el correo es: ", businessUser.Email)
		err = service.SignUp(&businessUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//generar el token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": businessUser.Email,
		})

		tokenString, errToken := token.SignedString(secretKey)
		if errToken != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
			return
		}

		response := gin.H{
			"token":        tokenString,
			"redirect_url": "http://localhost:5173/login", // Reemplaza con la URL correcta
		}

		c.JSON(http.StatusOK, response)
	}
}
