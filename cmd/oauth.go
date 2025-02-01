package cmd

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     "{}",
	ClientSecret: "{}",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www/googleapis.com/auth/photoslibrary.readonly",
	},
	Endpoint: google.Endpoint,
}
