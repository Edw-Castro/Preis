package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logout() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
