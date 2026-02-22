package auth

import (
	"context"
	"errors"
	"time"

	"github.com/khiz125/goapi/apperrors"
	"github.com/khiz125/goapi/common"
	"github.com/khiz125/goapi/domain/user"
)

func (s *AuthService) HandleGoogleCallback(
	ctx context.Context,
	code string,
) (*user.User, error) {

	googleUser, err := s.googleClient.ExchangeCodeForIDToken(ctx, code)

	if err != nil {
		return nil, apperrors.OAuthExchangeFailed.Wrap(err, "google oauth failed")
	}

	var result *user.User

	err = s.uow.Do(ctx, func(r user.Repositories) error {

		identity, err := r.Identity.FindByProviderSub("google", googleUser.Sub)
		if err == nil {
			result, err = r.User.FindByID(identity.UserID)
			return err
		}
		if !errors.Is(err, user.ErrIdentityNotFound) {
			return err
		}

		newUser := &user.User{
			ID:        common.GenerateUUID(),
			Name:      googleUser.Name,
			Email:     &googleUser.Email,
			CreatedAt: time.Now(),
		}

		if err := r.User.Create(newUser); err != nil {
			return err
		}

		newIdentity := &user.Identity{
			UserID:      newUser.ID,
			Provider:    "google",
			ProviderSub: googleUser.Sub,
			CreatedAt:   time.Now(),
		}

		if err := r.Identity.Create(newIdentity); err != nil {
			return err
		}

		result = newUser
		return nil

	})

	if err != nil {
		return nil, apperrors.GetDataFailed.Wrap(err, "google login failed")
	}

	return result, nil

}
