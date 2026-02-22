package user

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrIdentityNotFound = errors.New("identity not found")
)
