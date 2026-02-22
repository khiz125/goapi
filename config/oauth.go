package config

import "os"

const GoogleAuthURL = "https://accounts.google.com/o/oauth2/v2/auth"

type GoogleAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func LoadGoogleOAuthConfig() GoogleAuthConfig {
	return GoogleAuthConfig{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URI"),
	}
}
