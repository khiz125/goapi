package repositories

import (
	"database/sql"

	"github.com/khiz125/goapi/domain/user"
)

type UserIdentityRepository struct {
	db *sql.DB
}

func NewIdentityRepository(db *sql.DB) user.IdentityRepository {
	return &UserIdentityRepository{
		db: db,
	}
}

func (r *UserIdentityRepository) FindByProviderSub(provider, sub string) (*user.Identity, error) {
	query := `
  SELECT id, user_id, provider, provider_sub, created_at
  FROM user_identities
  WHERE provider = ? AND provider_sub = ?
  `

	row := r.db.QueryRow(query, provider, sub)

	var identity user.Identity

	err := row.Scan(
		&identity.ID,
		&identity.UserID,
		&identity.Provider,
		&identity.ProviderSub,
		&identity.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &identity, nil
}

func (r *UserIdentityRepository) Create(identity *user.Identity) error {

	query := `
  INSERT INTO user_identities
  (user_id, provider, provider_sub, created_at)
  VALUES (?, ?, ?, NOW())
  `

	_, err := r.db.Exec(
		query,
		identity.UserID,
		identity.Provider,
		identity.ProviderSub,
	)

	return err
}
