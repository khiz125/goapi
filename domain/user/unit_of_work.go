package user

import "context"

type UnitOfWork interface {
	Do(ctx context.Context, fn func(r Repositories) error) error
}
