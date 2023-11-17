package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/Edw-Castro/Preis/internal/auth/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GoogleAuthCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var userData map[string]interface{}
		code := c.Query("code")

		token, err := auth.GoogleOauthConfig.Exchange(c, code)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Guardar el token en la sesión
		session := sessions.Default(c)
		session.Set("access_token", token.AccessToken)
		session.Save()

		// Aquí puedes verificar la identidad del usuario y autorizarlo en tu aplicación.
		client := auth.GoogleOauthConfig.Client(c, token)
		userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer userInfo.Body.Close()

		err = json.NewDecoder(userInfo.Body).Decode(&userData)

		session.Set("email", userData["email"])
		session.Save()

		state := c.Query("state")
		fmt.Println("El estado es:", state)

		c.JSON(http.StatusFound, gin.H{"acces_token": token.AccessToken, "felicitacion": "MUY BIENNNN"})
	}

}

// //GPT
// return func(c *gin.Context) {
// 	code := c.Query("code")
// 	token, err := auth.GoogleOauthConfig.Exchange(c, code)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	client := auth.GoogleOauthConfig.Client(c, token)
// 	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer userInfo.Body.Close()

// 	var userData map[string]interface{}
// 	err = json.NewDecoder(userInfo.Body).Decode(&userData)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, userData)
// }
// //GPT
