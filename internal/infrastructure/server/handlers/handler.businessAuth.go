package server

import (
	"fmt"
	"net/http"

	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func User(service ports.BusinessOwnerService) func(*gin.Context) {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt-login")
		fmt.Println("Verificando si existe la cookie", cookie)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.SecretKeyClaims), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
			return
		}

		// Extraer las claims del token
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthenticated"})
			return
		}

		// Acceder a las claims
		userId := claims.Issuer
		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		fmt.Println("USER ID ES:", userId)
		// Puedes continuar con el código para buscar el usuario en la base de datos o realizar otras acciones según sea necesario.

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
