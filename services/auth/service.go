package auth

import "github.com/khiz125/goapi/domain/user"

type AuthService struct {
	userRepo     user.UserRepository
	identityRepo user.IdentityRepository
}

func NewAuthService(repo user.UserRepository, identity user.IdentityRepository) *AuthService {
	return &AuthService{
		userRepo:     repo,
		identityRepo: identity,
	}
}
