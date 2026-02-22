package auth

import (
	"github.com/khiz125/goapi/domain/user"
)

type AuthService struct {
	uow          user.UnitOfWork
	googleClient AuthGoogleClient
}

func NewAuthService(
	uow user.UnitOfWork,
	oauthClient AuthGoogleClient,
) *AuthService {
	return &AuthService{
		uow:          uow,
		googleClient: oauthClient,
	}
}
