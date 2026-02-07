package config

import "os"

type GoogleAuthConfig struct {
	AuthURL     string
	ClientID    string
	RedirectURI string
}

func LoadGoogleOAuthConfig() GoogleAuthConfig {
	return GoogleAuthConfig{
    AuthURL:      os.Getenv("GOOGLE_OAUTH_AUTH_URL"),
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		RedirectURI:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URI"),
	}
}
