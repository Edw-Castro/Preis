package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Edw-Castro/Preis/internal/core/domain"
	ports "github.com/Edw-Castro/Preis/internal/core/ports/services"
	"github.com/Edw-Castro/Preis/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login(service ports.BusinessOwnerService) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginBusiness domain.User
		fmt.Println("Si recibi el correo", loginBusiness.Email)
		if err := c.BindJSON(&loginBusiness); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userBusiness, err := service.Login(&loginBusiness)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email or password incorrect"})
			return
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(userBusiness.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		})

		token, err := claims.SignedString([]byte(utils.SecretKeyClaims))

		fmt.Println("eSTABLECIENDO COOKIES")

		// Imprimir información sobre la ubicación actual

		if err != nil {
			fmt.Println("Entre al error en el token")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Valor del token", token)
		// Establece la cookie en la respuesta

		c.SetCookie("jwt-login", token, int(time.Now().Add(24*time.Hour).Unix()), "/", "", false, false)
		c.JSON(http.StatusOK, "success")
	}
}
