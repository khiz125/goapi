package repositories

import (
	"context"
	"database/sql"

	"github.com/khiz125/goapi/domain/user"
)

type UnitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Do(ctx context.Context, fn func(user.Repositories) error) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	repos := user.Repositories{
		User:     NewUserRepository(tx),
		Identity: NewIdentityRepository(tx),
	}
	if err := fn(repos); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
