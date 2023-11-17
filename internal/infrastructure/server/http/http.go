package http

import (
	"net/http"

	"github.com/Edw-Castro/Preis/internal/infrastructure/server/middlewares"
	routes "github.com/Edw-Castro/Preis/internal/infrastructure/server/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/auth/login", routes.LoginGoogle())
	engine.GET("/auth/callback", routes.GoogleAuthCallback())
	engine.POST("/auth/register", routes.Register())
	engine.GET("product/:id", routes.ProductDetail())
	engine.POST("/business-owner", routes.SignUpBusinessOwner())
	engine.GET("/home", middlewares.AuthMiddleware(), func(c *gin.Context) {
		// Coloca aquí la lógica para proteger la ruta de inicio.
		c.JSON(http.StatusOK, gin.H{"message": "Esta es la página de inicio protegida"})
	})

}

/*engine.POST("/login", func(c *gin.Context) {
	var request struct {
		Name     string `form:"name" binding:"required"`
		Lastname string `form:"lastname" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("Email:", request.Email)
	fmt.Println("Password:", request.Password)
	fmt.Println("lastname", request.Lastname)
	fmt.Println("name", request.Name)
	// Realiza la autenticación y genera una respuesta, si es necesario
	var res struct {
		Message string `json:"message"`
	}
	res.Message = "Inicio de sesión exitoso"

	c.JSON(http.StatusOK, res)

})
*/
