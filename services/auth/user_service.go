package auth

import (
	"github.com/khiz125/goapi/common"
	"github.com/khiz125/goapi/domain/user"
)

func (s *AuthService) HandleGoogleCallback(code string) (*user.User, error) {

	// Todo get sub and other info from google

	provider := "google"
	sub := "sub from google"

	identity, err := s.identityRepo.FindByProviderSub(provider, sub)
	if err != nil {
		return s.userRepo.FindByID(identity.UserID)
	}

	newUser := &user.User{
		ID:   common.GenerateUUID(),
		Name: "name from google",
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	newIdentity := &user.Identity{
		UserID:      newUser.ID,
		Provider:    provider,
		ProviderSub: sub,
	}

	err = s.identityRepo.Create(newIdentity)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
