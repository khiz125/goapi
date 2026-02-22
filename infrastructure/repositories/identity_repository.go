package repositories

import (
	"database/sql"
	"github.com/khiz125/goapi/domain/user"
)

type IdentityRepository struct {
	tx *sql.Tx
}

func NewIdentityRepository(tx *sql.Tx) *IdentityRepository {
	return &IdentityRepository{tx: tx}
}

func (r *IdentityRepository) Create(i *user.Identity) error {
	query := `
	INSERT INTO user_identities (user_id, provider, provider_sub, created_at)
	VALUES (?, ?, ?, ?)
	`
	_, err := r.tx.Exec(query,
		i.UserID,
		i.Provider,
		i.ProviderSub,
		i.CreatedAt,
	)
	return err
}

func (r *IdentityRepository) FindByProviderSub(provider, sub string) (*user.Identity, error) {
	query := `
	SELECT user_id, provider, provider_sub, created_at
	FROM user_identities
	WHERE provider = ? AND provider_sub = ?
	`

	i := &user.Identity{}
	err := r.tx.QueryRow(query, provider, sub).Scan(
		&i.UserID,
		&i.Provider,
		&i.ProviderSub,
		&i.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, user.ErrIdentityNotFound
	}
	if err != nil {
		return nil, err
	}

	return i, nil
}
