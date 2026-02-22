package oauth

import (
	"context"
	"errors"

	"github.com/khiz125/goapi/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type GoogleClient struct {
	config *oauth2.Config
}

type GoogleUser struct {
	Sub   string
	Email string
	Name  string
}

func NewGoogleClient(cfg config.GoogleAuthConfig) *GoogleClient {

	oauthCfg := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURI,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
	}

	return &GoogleClient{
		config: oauthCfg,
	}
}

func (g *GoogleClient) ExchangeCodeForIDToken(ctx context.Context, code string) (*GoogleUser, error) {
	token, err := g.config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("missing id_token")
	}

	payload, err := idtoken.Validate(ctx, rawIDToken, g.config.ClientID)
	if err != nil {
		return nil, err
	}

	return &GoogleUser{
		Sub:   payload.Subject,
		Email: payload.Claims["email"].(string),
		Name:  payload.Claims["name"].(string),
	}, nil
}
