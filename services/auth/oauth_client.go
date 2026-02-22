package auth

import (
	"context"

	"github.com/khiz125/goapi/infrastructure/oauth"
)

type AuthGoogleClient interface {
	ExchangeCodeForIDToken(ctx context.Context, code string) (*oauth.GoogleUser, error)
}
