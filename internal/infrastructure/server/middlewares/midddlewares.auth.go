package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Verificar si el token está presente en la sesión o en una cookie
		// Obtener el token de acceso de la sesión
		session := sessions.Default(c)
		token := session.Get("access_token")

		if token == nil {
			// Si no hay token, redirigir al usuario a la página de inicio de sesión
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		// Verificar la validez del token (por ejemplo, la fecha de vencimiento)
		if !isValidToken(token.(string)) {
			// Si el token no es válido, redirigir al usuario a la página de inicio de sesión
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		fmt.Println("path url", c.Request.URL.Path)
		// Si el token es válido, permitir el acceso a la ruta protegida
		//c.JSON(http.StatusOK, gin.H{"token": token})
		c.Next()
	}

}

// Función para verificar si el token de acceso es válido
func isValidToken(accessToken string) bool {
	// Parsea el token (asumiendo que es un token JWT)
	token, err := parseAccessToken(accessToken)
	if err != nil {
		// Si hay un error al analizar el token, considerarlo no válido
		return false
	}

	// Verifica si la fecha de vencimiento del token aún no ha pasado
	currentTime := time.Now()
	return token.Expiry.After(currentTime)
}

// Función para analizar el token (puede variar según la implementación real)
func parseAccessToken(accessToken string) (*Token, error) {
	// Aquí deberías implementar la lógica para analizar el token JWT
	// y extraer la fecha de vencimiento y otros datos relevantes.

	// Por simplicidad, este es un ejemplo ficticio de cómo podría funcionar.
	// Debes reemplazarlo con una implementación adecuada.
	token := &Token{
		Expiry: time.Now().Add(time.Hour), // Fecha de vencimiento ficticia
	}

	return token, nil
}

// Estructura de token (puede variar según la implementación real)
type Token struct {
	Expiry time.Time
}
