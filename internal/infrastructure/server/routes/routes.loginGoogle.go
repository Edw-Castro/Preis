package routes

import (
	"fmt"
	"net/http"

	auth "github.com/Edw-Castro/Preis/internal/auth/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func LoginGoogle() func(c *gin.Context) {
	return func(c *gin.Context) {
		originURL := c.Query("originURL")
		fmt.Println("origingURL de LoginGoogle es", originURL)
		url := auth.GoogleOauthConfig.AuthCodeURL(originURL, oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("redirect_uri", "http://localhost:8080/auth/callback"))
		c.Redirect(http.StatusFound, url)
	}
}
