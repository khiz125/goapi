package repositories

import (
	"database/sql"
	"github.com/khiz125/goapi/domain/user"
)

type Repositories struct {
	userRepo     *UserRepository
	identityRepo *IdentityRepository
}

func NewRepositories(tx *sql.Tx) *Repositories {
	return &Repositories{
		userRepo:     NewUserRepository(tx),
		identityRepo: NewIdentityRepository(tx),
	}
}

func (r *Repositories) User() user.UserRepository {
	return r.userRepo
}

func (r *Repositories) Identity() user.IdentityRepository {
	return r.identityRepo
}
