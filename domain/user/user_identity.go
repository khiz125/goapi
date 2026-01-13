package user

import "time"

type Identity struct {
	ID          int64
	UserID      string
	Provider    string
	ProviderSub string
	CreatedAt   time.Time
}
