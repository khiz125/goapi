package repositories

import (
	"database/sql"
	"github.com/khiz125/goapi/domain/user"
)

type UserRepository struct {
	tx *sql.Tx
}

func NewUserRepository(tx *sql.Tx) *UserRepository {
	return &UserRepository{tx: tx}
}

func (r *UserRepository) Create(u *user.User) error {

	query := `
  INSERT into users (id, name, email, created_at)
  VALUES (?,?,?,?)
  `

	_, err := r.tx.Exec(
		query,
		u.ID,
		u.Name,
		u.Email,
		u.CreatedAt,
	)

	return err
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
	query := `
  SELECT id, name, email, created_at
  FROM users
  WHERE id = ?
  `

	u := &user.User{}

	err := r.tx.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByGoogleSub(sub string) (*user.User, error) {

	query := `
  SELECT u.id, u.name, u.email, u. created_at
  FROM users u
  JOIN user_identities ui
    ON u.id = ui.user_id
  WHERE ui.provider = 'google'
    AND ui.provider_sub = ?
  `

	u := &user.User{}

	err := r.tx.QueryRow(query, sub).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
