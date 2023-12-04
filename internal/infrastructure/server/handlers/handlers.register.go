package server

import (
	"fmt"
	"net/http"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Clave secreta para firmar tokens JWT (mejorarlo para obtnerlo desde env)

func Register(service ports.BusinessOwnerService) func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Comprobando si existen mail, name", c.Query("email"), c.Query("name"))
		var data map[string]string

		if err := c.Bind(&data); err != nil {
			fmt.Println("Tengo un error con la data")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		pwd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
		// fmt.Println("nombre", registerUser.Name, "apellido", registerUser.LastName, "email", registerUser.Email, "password: ", registerUser.Password)
		// fmt.Println(registerUser.Email)
		businessUser := domain.User{
			Email:    data["email"],
			Password: string(pwd),
			// Asigna otros valores según sea necesario
			Name: data["name"],
			Role: "customer",
		}
		fmt.Println("el correo es: ", businessUser.Email)

		err := service.SignUp(&businessUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, businessUser)
	}
}
