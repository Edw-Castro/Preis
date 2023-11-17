package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	ClientID:     "382481055888-lbuhf2m0g0caulrpfioven5ur8a8h3v2.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/callback",
	Endpoint:     google.Endpoint,
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.email", "openid",
	},
}
