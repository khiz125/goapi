package user

import "time"

type Credential struct {
	UserID       string
	PasswordHash string
	CreatedAt    time.Time
}
