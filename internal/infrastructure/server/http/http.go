package http

import (
	routes "github.com/Edw-Castro/Preis/internal/infrastructure/server/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/preis/google/auth/login", routes.LoginGoogle())
	engine.GET("/preis/google/auth/callback", routes.GoogleAuthCallback())
	engine.GET("/preis/auth/user", routes.User())
	engine.POST("/preis/auth/register", routes.Register())
	engine.POST("/preis/auth/login", routes.Login())
	engine.POST("/preis/auth/logout", routes.Logout())
	engine.GET("/preis/product/:id", routes.ProductDetail())
	engine.GET("/preis/product/upload", routes.ProductUpload())
	engine.POST("/preis/product/compare", routes.ProductCompare())
	engine.POST("/preis/business-owner", routes.SignUpBusinessOwner())

}
